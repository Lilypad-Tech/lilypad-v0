package bridge

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"go.ptx.dk/multierrgroup"
	"go.uber.org/multierr"
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
}

// Run the workflow: listen to the smart contract and process any orders that
// come through.
//
// Run accepts two contexts. One will be used for the "plumbing" (components
// that persist, retry, etc) and the other for the "actors" (components that
// communicate with external services). If the actor context is cancelled,
// actors will be allowed to finish their current operation before exiting. If
// the plumbing context is cancelled, all actor operation will cease
// immediately.
//
// Run will blocks until the plumbing context or actor context are cancelled.
func (workflow *Workflow) Run(plumbCtx, actorCtx context.Context, plumbingCancel context.CancelFunc) (err error) {
	var actors, plumbing multierrgroup.Group

	// Listen to the smart contract for any submitted orders.
	jobRequests := make(chan ContractSubmittedEvent, 1)
	defer close(jobRequests)

	actors.Go(func() error {
		return workflow.Contract.Listen(actorCtx, jobRequests)
	})

	jobRequestsSaved := make(chan ContractSubmittedEvent, 256)
	defer close(jobRequestsSaved)

	plumbing.Go(func() error {
		Persist(plumbCtx, jobRequests, jobRequestsSaved)
		return nil
	})

	// If we fail to do something too many times, it's possible that the user's
	// job is just unrunnable. So we will refund their job in that case.
	jobsToRefund := make(chan ContractSubmittedEvent, 256)
	defer close(jobsToRefund)

	jobsRefunded := make(chan ContractRefundedEvent, 256)
	defer close(jobsRefunded)

	actors.Go(func() error {
		ErrorActor(
			log.Ctx(plumbCtx).With().Str("action", "Refund").Int("instance", 0).Logger().WithContext(plumbCtx),
			log.Ctx(actorCtx).With().Str("action", "Refund").Int("instance", 0).Logger().WithContext(actorCtx),
			jobsToRefund,
			workflow.Contract.Refund,
			jobsRefunded,
			jobsToRefund, // TODO: apply retry back-off
		)
		return nil
	})

	// When a order is submitted, kick-off a Bacalhau job to run it. We use
	// three actors so that we can handle lots of contracts coming in at once.
	// If we fail to create a job, retry it up to 5 times before giving up.
	jobsInProgress := make(chan BacalhauJobRunningEvent, 256)
	defer close(jobsInProgress)

	jobsToRetryCreate := make(chan ContractSubmittedEvent, 256)
	defer close(jobsToRetryCreate)

	for i := 3; i > 0; i-- {
		inst := i
		actors.Go(func() error {
			ErrorActor(
				log.Ctx(plumbCtx).With().Str("action", "CreateJob").Int("instance", inst).Logger().WithContext(plumbCtx),
				log.Ctx(actorCtx).With().Str("action", "CreateJob").Int("instance", inst).Logger().WithContext(actorCtx),
				jobRequestsSaved,
				workflow.Bacalhau.Create,
				jobsInProgress,
				jobsToRetryCreate,
			)
			return nil
		})
	}

	plumbing.Go(func() error {
		ctx := log.Ctx(plumbCtx).With().Str("action", "Retry:CreateJob").Logger().WithContext(plumbCtx)
		Retry(ctx, 5, jobsToRetryCreate, jobRequestsSaved, jobsToRefund)
		return nil
	})

	jobsInProgressSaved := make(chan BacalhauJobRunningEvent, 1)
	defer close(jobsInProgressSaved)

	plumbing.Go(func() error {
		Persist(plumbCtx, jobsInProgress, jobsInProgressSaved)
		return nil
	})

	// For now, until we have some persistence, just pass the job we created
	// straight onto the completion stage. Once we have persistence, we can just
	// load the current list of jobs every time.
	jobBatchesInProgress := make(chan []BacalhauJobRunningEvent, 256)
	defer close(jobBatchesInProgress)

	plumbing.Go(func() error {
		Actor(
			log.Ctx(plumbCtx).With().Str("action", "FetchJob").Int("instance", 0).Logger().WithContext(plumbCtx),
			log.Ctx(actorCtx).With().Str("action", "FetchJob").Int("instance", 0).Logger().WithContext(actorCtx),
			jobsInProgressSaved,
			func(ctx context.Context, bjre BacalhauJobRunningEvent) []BacalhauJobRunningEvent {
				return []BacalhauJobRunningEvent{bjre}
			},
			jobBatchesInProgress,
		)
		return nil
	})

	// Every five seconds, check that status of all of our Bacalhau jobs. Doing
	// it this way allows the system to be minimise the number of Bacalhau API
	// calls. For any jobs that have failed, try to run them again.
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.SetMaxConcurrentJobs(1, gocron.WaitMode)
	scheduler.StartAsync()
	defer scheduler.Stop()

	// fetchJobs := func() { Fetch(actorCtx, jobBatchesInProgress) }
	// _, err = scheduler.Every(5).Seconds().Do(fetchJobs)
	// if err != nil {
	// 	return err
	// }

	jobBatchesCompleted := make(chan []BacalhauJobCompletedEvent, 256)
	defer close(jobBatchesCompleted)

	jobBatchesFailed := make(chan []BacalhauJobFailedEvent, 16)
	defer close(jobBatchesFailed)

	actors.Go(func() error {
		TwoActor(
			log.Ctx(actorCtx).With().Str("action", "FindCompleted").Int("instance", 0).Logger().WithContext(actorCtx),
			log.Ctx(plumbCtx).With().Str("action", "FindCompleted").Int("instance", 0).Logger().WithContext(plumbCtx),
			jobBatchesInProgress,
			workflow.Bacalhau.FindCompleted,
			jobBatchesCompleted,
			jobBatchesFailed,
		)
		return nil
	})

	jobsCompleted := make(chan BacalhauJobCompletedEvent, 256)
	defer close(jobsCompleted)

	plumbing.Go(func() error {
		Flatten(plumbCtx, jobBatchesCompleted, jobsCompleted)
		return nil
	})

	jobsFailed := make(chan BacalhauJobFailedEvent, 256)
	defer close(jobsFailed)

	plumbing.Go(func() error {
		Flatten(plumbCtx, jobBatchesFailed, jobsFailed)
		return nil
	})

	retryableJobs := make(chan BacalhauJobFailedEvent)
	plumbing.Go(func() error {
		Actor(plumbCtx, actorCtx, retryableJobs, func(ctx context.Context, e BacalhauJobFailedEvent) ContractSubmittedEvent {
			return e
		}, jobRequests)
		return nil
	})

	refundableJobs := make(chan BacalhauJobFailedEvent)
	plumbing.Go(func() error {
		Actor(plumbCtx, actorCtx, refundableJobs, func(ctx context.Context, e BacalhauJobFailedEvent) ContractSubmittedEvent {
			return e
		}, jobsToRefund)
		return nil
	})

	plumbing.Go(func() error {
		ctx := log.Ctx(plumbCtx).With().Str("action", "Retry:FindCompleted").Logger().WithContext(plumbCtx)
		Retry(ctx, 3, jobsFailed, retryableJobs, refundableJobs)
		return nil
	})

	// For jobs that have completed, send them back to the smart contract for payment.
	jobsPaid := make(chan ContractPaidEvent, 1)
	defer close(jobsPaid)

	jobsFailedToPay := make(chan BacalhauJobCompletedEvent, 256)
	defer close(jobsFailedToPay)

	actors.Go(func() error {
		ErrorActor(
			log.Ctx(actorCtx).With().Str("action", "Payment").Int("instance", 0).Logger().WithContext(actorCtx),
			log.Ctx(plumbCtx).With().Str("action", "Payment").Int("instance", 0).Logger().WithContext(plumbCtx),
			jobsCompleted,
			workflow.Contract.Complete,
			jobsPaid,
			jobsFailedToPay,
		)
		return nil
	})

	jobsFailedToPayToRefund := make(chan BacalhauJobCompletedEvent, 1)
	plumbing.Go(func() error {
		Actor(plumbCtx, actorCtx, jobsFailedToPayToRefund, func(ctx context.Context, e BacalhauJobCompletedEvent) ContractSubmittedEvent {
			return e
		}, jobsToRefund)
		return nil
	})

	plumbing.Go(func() error {
		Retry(plumbCtx, 20, jobsFailedToPay, jobsCompleted, jobsFailedToPayToRefund)
		return nil
	})

	jobsPaidSaved := make(chan ContractPaidEvent, 256)
	defer close(jobsPaidSaved)

	plumbing.Go(func() error {
		Persist(plumbCtx, jobsPaid, jobsPaidSaved)
		return nil
	})

	plumbing.Go(func() error {
		Discard(plumbCtx, jobsPaidSaved)
		return nil
	})

	log.Ctx(actorCtx).Info().Msg("Ready")
	defer log.Ctx(actorCtx).Info().Msg("Shutdown")

	// Sit and wait on the actors. If they exit, then the graceful shutdown
	// context has been triggered, so now shutdown the plumbing.
	err = actors.Wait()
	plumbingCancel()
	return multierr.Combine(err, plumbing.Wait())
}
