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
// The Workflow is modelled as a set of Actors – goroutines that read from and
// post to channels. This model allows us to encapsulate the complex logic of
// dealing with a smart contract and Bacalhau into a simple components, but then
// to scale these components efficiently as we receive more and more load on the
// service.
//
// For example, using this model means we can have three workers posting to the
// Bacalhau network at once – something that is is relatively slow and would be
// a bottleneck at high load.
//
// It also allows Actors to be composed with generic logic that handles
// persistence and retries, so that we can easily handle failing to perform some
// action without having to pollute those components with the detail of how
// those retries or persists are done.
type Workflow struct {
	Bacalhau JobRunner
	Contract SmartContract
	Repo     Repository

	scheduler *gocron.Scheduler
}

func NewWorkflow(jr JobRunner, sc SmartContract, repo Repository) *Workflow {
	return &Workflow{
		Bacalhau:  jr,
		Contract:  sc,
		Repo:      repo,
		scheduler: gocron.NewScheduler(time.UTC),
	}
}

var maxAttemptsByState map[OrderState]int = map[OrderState]int{
	OrderStateSubmitted: 5,
	OrderStateRunning:   0,
	OrderStateCompleted: 5,
	OrderStateJobError:  3,
	OrderStateFailed:    5,
	OrderStatePaid:      0,
	OrderStateRefunded:  0,
}

var jobCheckInterval time.Duration = 5 * time.Second

func (workflow *Workflow) Start(ctx context.Context) error {
	wg := multierrgroup.Group{}

	submittedEvents := make(chan ContractSubmittedEvent)
	defer close(submittedEvents)
	newEvents := make(chan Event, 256)
	defer close(newEvents)

	wg.Go(func() error { return workflow.Run(ctx, newEvents) })
	wg.Go(func() error { return workflow.Contract.Listen(ctx, submittedEvents) })

	go func() {
		for {
			select {
			case e := <-submittedEvents:
				newEvents <- e
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Go(func() error {
		workflow.scheduler.StartAsync()
		<-ctx.Done()
		workflow.scheduler.Stop()
		workflow.scheduler.Clear()
		return nil
	})

	_, err := workflow.scheduler.Every(jobCheckInterval).Do(func() {
		workflow.CheckRunningEvents(ctx, newEvents)
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

	return wg.Wait()
}

func (workflow *Workflow) Run(ctx context.Context, newEvents <-chan Event) error {
	processedEvents := make(chan Event, 256)

	for {
		var result Event
		select {
		case event := <-processedEvents:
			result = workflow.ProcessEvent(ctx, event)
		case event := <-newEvents:
			result = workflow.ProcessEvent(ctx, event)
		case <-ctx.Done():
			return nil
		}

		if result != nil {
			processedEvents <- result
		}
	}
}

func (workflow *Workflow) ProcessEvent(ctx context.Context, event Event) (result Event) {
	var err error
	log.Ctx(ctx).Trace().Stringer("id", event.OrderId()).Stringer("state", event.OrderState()).Msg("Process event")

	currentState := event.OrderState()
	switch currentState {
	case OrderStateSubmitted:
		result, err = workflow.Bacalhau.Create(ctx, event.(ContractSubmittedEvent))
	case OrderStateCompleted:
		result, err = workflow.Contract.Complete(ctx, event.(BacalhauJobCompletedEvent))
	case OrderStateJobError:
		event := event.(BacalhauJobFailedEvent)
		if workflow.ShouldRetry(event) {
			result = event.Retry()
		} else {
			result = event.Failed()
		}
	case OrderStateFailed:
		result, err = workflow.Contract.Refund(ctx, event.(ContractFailedEvent))
	default:
		result = nil
	}

	if err != nil && !errors.Is(err, context.Canceled) {
		// The processing action failed. If we can retry the action, do that,
		// else if we are beyond our limit send the order for a refund.
		if e, retryable := event.(Retryable); retryable && workflow.ShouldRetry(e) {
			e.AddAttempt()
			result = e // TODO: retry back off
		} else {
			result = event.(ContractSubmittedEvent).Failed()
		}
	}

	// If we have a non-nil result, we are changing the state of something. So
	// save the new state.
	if result != nil {
		saveError := workflow.Repo.Save(result)
		log.Ctx(ctx).WithLevel(level(saveError)).
			Err(saveError).
			Stringer("oldState", currentState).
			Stringer("newState", result.OrderState()).
			Msg("Saving result")
	}

	return
}

func (workflow *Workflow) CheckRunningEvents(ctx context.Context, out chan<- Event) {
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

func (w *Workflow) ShouldRetry(event Retryable) bool {
	return event.Attempts() < maxAttemptsByState[event.OrderState()]
}

func level(err error) zerolog.Level {
	if err == nil {
		return zerolog.DebugLevel
	} else {
		return zerolog.ErrorLevel
	}
}
