package bridge

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"go.ptx.dk/multierrgroup"
	"go.uber.org/multierr"
)

func Run(plumbCtx, actorCtx context.Context, plumbingCancel context.CancelFunc) (err error) {
	var actors, plumbing multierrgroup.Group

	jobRunner := NewJobRunner()
	smartContract := MockContract()

	jobRequests := make(chan ContractSubmittedEvent, 1)
	defer close(jobRequests)

	actors.Go(func() error {
		return smartContract.Listen(actorCtx, jobRequests)
	})

	jobsToRefund := make(chan ContractSubmittedEvent, 256)
	defer close(jobsToRefund)

	jobsRefunded := make(chan ContractRefundedEvent, 256)
	defer close(jobsRefunded)

	actors.Go(func() error {
		ErrorActor(
			log.Ctx(plumbCtx).With().Str("action", "Refund").Int("instance", 0).Logger().WithContext(plumbCtx),
			log.Ctx(actorCtx).With().Str("action", "Refund").Int("instance", 0).Logger().WithContext(actorCtx),
			jobsToRefund,
			smartContract.Refund,
			jobsRefunded,
			jobsToRefund,
		)
		return nil
	})

	jobRequestsSaved := make(chan ContractSubmittedEvent, 256)
	defer close(jobRequestsSaved)

	plumbing.Go(func() error {
		Persist(plumbCtx, jobRequests, jobRequestsSaved)
		return nil
	})

	jobsInProgress := make(chan BacalhauJobRunningEvent, 256)
	defer close(jobsInProgress)

	jobsToRetryCreate := make(chan ContractSubmittedEvent, 256)
	defer close(jobsToRetryCreate)

	for i := 3; i > 0; i-- {
		inst := i
		actors.Go(func() error {
			ErrorActor(
				log.Ctx(actorCtx).With().Str("action", "CreateJob").Int("instance", inst).Logger().WithContext(actorCtx),
				log.Ctx(plumbCtx).With().Str("action", "CreateJob").Int("instance", inst).Logger().WithContext(plumbCtx),
				jobRequestsSaved,
				jobRunner.Create,
				jobsInProgress,
				jobsToRetryCreate,
			)
			return nil
		})
	}

	plumbing.Go(func() error {
		Retry(plumbCtx, 1, jobsToRetryCreate, jobRequestsSaved, jobsToRefund)
		return nil
	})

	jobsInProgressSaved := make(chan BacalhauJobRunningEvent, 1)
	defer close(jobsInProgressSaved)
	plumbing.Go(func() error {
		Persist(plumbCtx, jobsInProgress, jobsInProgressSaved)
		return nil
	})
	plumbing.Go(func() error {
		Discard(plumbCtx, jobsInProgressSaved)
		return nil
	})

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.SetMaxConcurrentJobs(1, gocron.WaitMode)
	scheduler.StartAsync()
	defer scheduler.Stop()

	jobBatchesInProgress := make(chan []BacalhauJobRunningEvent, 256)
	defer close(jobBatchesInProgress)
	fetchJobs := func() { Fetch(actorCtx, jobBatchesInProgress) }
	_, err = scheduler.Every(5).Seconds().Do(fetchJobs)
	if err != nil {
		return err
	}

	jobBatchesCompleted := make(chan []BacalhauJobCompletedEvent, 256)
	defer close(jobBatchesCompleted)

	jobBatchesFailed := make(chan []BacalhauJobFailedEvent, 16)
	defer close(jobBatchesFailed)

	actors.Go(func() error {
		TwoActor(
			log.Ctx(actorCtx).With().Str("action", "FindCompleted").Int("instance", 0).Logger().WithContext(actorCtx),
			log.Ctx(plumbCtx).With().Str("action", "FindCompleted").Int("instance", 0).Logger().WithContext(plumbCtx),
			jobBatchesInProgress,
			jobRunner.FindCompleted,
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
		Retry(plumbCtx, 3, jobsFailed, retryableJobs, refundableJobs)
		return nil
	})

	log.Ctx(actorCtx).Info().Msg("Ready")
	defer log.Ctx(actorCtx).Info().Msg("Shutdown")

	err = actors.Wait()
	plumbingCancel()
	return multierr.Combine(err, plumbing.Wait())
}
