package bridge

import (
	"context"

	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/ipfs/go-cid"
	"github.com/pkg/errors"
)

type RunnerCreateHandler func(context.Context, ContractSubmittedEvent) (BacalhauJobRunningEvent, error)
type RunnerFindCompletedHandler func(context.Context, []BacalhauJobRunningEvent) ([]BacalhauJobCompletedEvent, []BacalhauJobFailedEvent)

var SuccessfulCreate RunnerCreateHandler = func(ctx context.Context, cse ContractSubmittedEvent) (BacalhauJobRunningEvent, error) {
	return cse.JobCreated(model.NewJob()), nil
}

var ErrorCreate RunnerCreateHandler = func(ctx context.Context, cse ContractSubmittedEvent) (BacalhauJobRunningEvent, error) {
	return nil, errors.New("error creating job")
}

var SuccssfulFind RunnerFindCompletedHandler = func(ctx context.Context, jobs []BacalhauJobRunningEvent) ([]BacalhauJobCompletedEvent, []BacalhauJobFailedEvent) {
	completed := []BacalhauJobCompletedEvent{}
	for _, job := range jobs {
		completed = append(completed, job.Completed(cid.Cid{}, "", "", 0))
	}
	return completed, nil
}

var FailedFind RunnerFindCompletedHandler = func(ctx context.Context, jobs []BacalhauJobRunningEvent) ([]BacalhauJobCompletedEvent, []BacalhauJobFailedEvent) {
	failed := []BacalhauJobFailedEvent{}
	for _, job := range jobs {
		failed = append(failed, job.JobError(""))
	}
	return nil, failed
}

// A JobRunner that won't make real requests and instead just runs the supplied
// functions when its methods are called.
type mockRunner struct {
	CreateHandler        RunnerCreateHandler
	FindCompletedHandler RunnerFindCompletedHandler
}

// Create implements JobRunner
func (mock *mockRunner) Create(ctx context.Context, job ContractSubmittedEvent) (BacalhauJobRunningEvent, error) {
	if mock.CreateHandler != nil {
		return mock.CreateHandler(ctx, job)
	} else {
		return SuccessfulCreate(ctx, job)
	}
}

// FindCompleted implements JobRunner
func (mock *mockRunner) FindCompleted(ctx context.Context, jobs []BacalhauJobRunningEvent) ([]BacalhauJobCompletedEvent, []BacalhauJobFailedEvent) {
	if mock.FindCompletedHandler != nil {
		return mock.FindCompletedHandler(ctx, jobs)
	} else {
		return SuccssfulFind(ctx, jobs)
	}
}

var _ JobRunner = (*mockRunner)(nil)
