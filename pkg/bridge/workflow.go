package bridge

import (
	"context"
	"errors"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.ptx.dk/multierrgroup"
)

// A Workflow is an active component that runs the entire process of an order
// being made on a smart contract, a job being run on Bacalhau, and the smart
// contract being paid.
//
// The Workflow is modelled as a simple state machine that hands off state
// transitions to other objects. This model allows us to encapsulate the complex
// logic of dealing with a smart contract and Bacalhau into simple components,
// but then to apply common persistance and retry logic to all states.
//
// Note that the workflow does not include an action to check that an individual
// job is running on Bacalhau. Because jobs are relatively long lived (~minutes)
// we would quickly fill up our buffers with jobs in the running state and have
// no room for any more new ones. So as an optimisation, we instead periodically
// reload all of the running jobs from the DB and check all of them at once, and
// only post them back into the state machine if they are finished.
//
// Note also that the workflow uses two work queues, one for adding events to
// the state machine and one for continuing to process those already present.
// Pre-existing events are processed over new ones. This is to ensure that new
// events can't clog up the system to such an extent that the workers can't push
// results onto the queue anymore, and thus would be deadlocked.
type Workflow struct {
	Bacalhau JobRunner
	Contract SmartContract
	Repo     Repository

	scheduler        *gocron.Scheduler
	getRetryTime     RetryStrategy
	jobCheckInterval time.Duration
}

var (
	defaultJobCheckInterval time.Duration = 5 * time.Second
	defaultRetryStrategy    RetryStrategy = Exponential
)

func NewWorkflow(jr JobRunner, sc SmartContract, repo Repository) *Workflow {
	return &Workflow{
		Bacalhau:         jr,
		Contract:         sc,
		Repo:             repo,
		scheduler:        gocron.NewScheduler(time.UTC),
		getRetryTime:     defaultRetryStrategy,
		jobCheckInterval: defaultJobCheckInterval,
	}
}

// Start spins up all of the goroutines that will generate and process items in
// the workflow. It will block until the passed context is cancelled.
func (workflow *Workflow) Start(ctx context.Context) error {
	wg := multierrgroup.Group{}

	submittedEvents := make(chan ContractSubmittedEvent)
	defer close(submittedEvents)
	newEvents := make(chan Event, 256)
	defer close(newEvents)

	wg.Go(func() error { return workflow.Run(ctx, newEvents) })
	wg.Go(func() error { return workflow.Contract.Listen(ctx, submittedEvents) })
	wg.Go(func() error { return workflow.deduplicateSubmittedEvents(ctx, submittedEvents, newEvents) })
	wg.Go(func() error {
		workflow.scheduler.StartAsync()
		<-ctx.Done()
		workflow.scheduler.Stop()
		workflow.scheduler.Clear()
		return nil
	})

	_, err := workflow.scheduler.Every(workflow.jobCheckInterval).Do(func() {
		workflow.checkRunningEvents(ctx, newEvents)
	})
	if err != nil {
		return err
	}

	wg.Go(func() error {
		return ReloadToChan[ContractSubmittedEvent](workflow.Repo, OrderStateSubmitted, newEvents)
	})
	wg.Go(func() error {
		return ReloadToChan[ContractFailedEvent](workflow.Repo, OrderStateFailed, newEvents)
	})
	wg.Go(func() error {
		return ReloadToChan[BacalhauJobCompletedEvent](workflow.Repo, OrderStateCompleted, newEvents)
	})
	wg.Go(func() error {
		return ReloadToChan[BacalhauJobFailedEvent](workflow.Repo, OrderStateJobError, newEvents)
	})

	log.Ctx(ctx).Info().Msg("Bridge ready")
	defer log.Ctx(ctx).Info().Msg("Bridge shutdown")

	return wg.Wait()
}

// Run processes events on the work queue, transitioning them through the state
// machine. It will block until the passed context is cancelled.
func (workflow *Workflow) Run(ctx context.Context, newEvents <-chan Event) (err error) {
	processedEvents := make(chan Event, 256)

	for {
		var result Event
		var wait time.Duration
		select {
		case event := <-processedEvents:
			result, wait = workflow.ProcessEvent(ctx, event)
		case event := <-newEvents:
			result, wait = workflow.ProcessEvent(ctx, event)
		case <-ctx.Done():
			return
		}

		if result != nil && wait == 0 {
			// We got an event result, so put it back on the queue to be
			// processed into the next state.
			processedEvents <- result
		} else if result != nil && wait > 0 {
			// The result has come with a wait time. Ask the scheduler to put
			// the result back on the queue after the wait time has elapsed.
			_, err = workflow.scheduler.WaitForSchedule().
				Every(1).
				LimitRunsTo(1).
				StartAt(time.Now().Add(wait)).
				Do(func() {
					processedEvents <- result
				})
		}

		if err != nil {
			log.Ctx(ctx).Error().Err(err).Send()
		}
	}
}

// ProcessEvent will transition a single event through the state machine,
// returning the event in a new state (or the same state if the action failed).
//
// If the returned event is nil, then the event shouldn't be processed any more.
// If the resturned wait time is non-zero, further actions on that event should
// be delayed until that time has elapsed.
func (workflow *Workflow) ProcessEvent(ctx context.Context, event Event) (result Event, wait time.Duration) {
	var err error
	ctx = log.Ctx(ctx).With().
		Stringer("id", event.OrderId()).
		Stringer("state", event.OrderState()).
		Logger().WithContext(ctx)

	log.Ctx(ctx).Trace().Msg("Process event")

	currentState := event.OrderState()
	switch currentState {
	case OrderStateSubmitted:
		result, err = workflow.Bacalhau.Create(ctx, event.(ContractSubmittedEvent))
	case OrderStateCompleted:
		result, err = workflow.Contract.Complete(ctx, event.(BacalhauJobCompletedEvent))
	case OrderStateJobError:
		event := event.(BacalhauJobFailedEvent)
		if ShouldRetry(event) {
			result = event.Retry()
		} else {
			result = event.Failed(event.Error())
		}
	case OrderStateFailed:
		result, err = workflow.Contract.Refund(ctx, event.(ContractFailedEvent))
	default:
		result = nil
	}

	if err != nil && !errors.Is(err, context.Canceled) {
		log.Ctx(ctx).Error().Err(err).Msg("Error processing event")

		// The processing action failed. If we can retry the action, do that,
		// else if we are beyond our limit send the order for a refund.
		if e, retryable := event.(Retryable); retryable && ShouldRetry(e) {
			e.AddAttempt()
			result = e
			wait = workflow.getRetryTime(e)
		} else {
			result = event.(ContractSubmittedEvent).Failed(err.Error())
		}
	}

	// If we have a non-nil result, we are changing the state of something. So
	// save the new state.
	if result != nil {
		saveError := workflow.Repo.Save(result)
		log.Ctx(ctx).WithLevel(level(saveError)).
			Err(saveError).
			Stringer("old", currentState).
			Stringer("new", result.OrderState()).
			Msg("Saving result")
	}

	return
}

// checkRunningEvents reloads the Bacalhau jobs that should be running and finds
// the ones that have finished, pushing them back onto the state machine queue
// for the result of the job to be processed.
func (workflow *Workflow) checkRunningEvents(ctx context.Context, out chan<- Event) {
	jobs, err := Reload[BacalhauJobRunningEvent](workflow.Repo, OrderStateRunning)
	log.Ctx(ctx).WithLevel(level(err)).Err(err).Int("count", len(jobs)).Msg("Reloaded running events")

	completed, failed := workflow.Bacalhau.FindCompleted(ctx, jobs)
	log.Ctx(ctx).Debug().
		Int("completed", len(completed)).
		Int("failed", len(failed)).
		Msg("Queried Bacalhau job status")

	for _, event := range completed {
		out <- event
	}
	for _, event := range failed {
		out <- event
	}
}

func level(err error) zerolog.Level {
	if err == nil {
		return zerolog.DebugLevel
	} else {
		return zerolog.ErrorLevel
	}
}

func (workflow *Workflow) deduplicateSubmittedEvents(ctx context.Context, in <-chan ContractSubmittedEvent, out chan<- Event) (err error) {
	for {
		select {
		case e := <-in:
			var exists bool
			exists, err = workflow.Repo.Exists(e)
			if err != nil {
				break
			}
			if exists {
				log.Ctx(ctx).Debug().Stringer("id", e.OrderId()).Msg("Dropping new event because already seen")
				continue
			}

			err = workflow.Repo.Save(e)
			out <- e
		case <-ctx.Done():
			return
		}

		if err != nil {
			log.Ctx(ctx).Error().Err(err)
		}
	}
}
