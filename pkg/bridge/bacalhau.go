package bridge

import (
	"context"
	"fmt"

	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/filecoin-project/bacalhau/pkg/publicapi"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func init() {
	err := system.InitConfig()
	if err != nil {
		panic(err)
	}
}

type JobRunner interface {
	Create(ctx context.Context, job ContractSubmittedEvent) (BacalhauJobRunningEvent, error)

	FindCompleted(ctx context.Context, jobs []BacalhauJobRunningEvent) ([]BacalhauJobCompletedEvent, []BacalhauJobFailedEvent)
}

type xRunner struct {
	Client *publicapi.APIClient
}

// Create implements JobRunner
func (r *xRunner) Create(ctx context.Context, e ContractSubmittedEvent) (BacalhauJobRunningEvent, error) {
	job, err := model.NewJobWithSaneProductionDefaults()
	if err != nil {
		return nil, errors.Wrap(err, "error creating Bacalhau job")
	}

	job.Spec = e.Spec()
	job, err = r.Client.Submit(ctx, job, nil)
	if err != nil {
		err = errors.Wrap(err, "error submitting Bacalhau job")
	}

	log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Str("job", job.Metadata.ID).Msg("Created Bacalhau job")
	return e.JobCreated(job), err
}

// FindCompleted implements JobRunner
func (*xRunner) FindCompleted(ctx context.Context, jobs []BacalhauJobRunningEvent) ([]BacalhauJobCompletedEvent, []BacalhauJobFailedEvent) {
	panic("unimplemented")
}

var _ JobRunner = (*xRunner)(nil)

func NewJobRunner() JobRunner {
	apiPort := 1234
	apiHost := "bootstrap.production.bacalhau.org"
	client := publicapi.NewAPIClient(fmt.Sprintf("http://%s:%d", apiHost, apiPort))
	return &xRunner{Client: client}
}
