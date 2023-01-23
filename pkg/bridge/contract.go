package bridge

import (
	"context"
	"time"

	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type SmartContract interface {
	Listen(context.Context, chan<- ContractSubmittedEvent) error

	Complete(context.Context, BacalhauJobCompletedEvent) (ContractPaidEvent, error)

	Refund(context.Context, ContractSubmittedEvent) (ContractRefundedEvent, error)
}

type mockContract struct{}

// Complete implements SmartContract
func (mockContract) Complete(ctx context.Context, e BacalhauJobCompletedEvent) (ContractPaidEvent, error) {
	log.Ctx(ctx).Debug().Stringer("id", e.OrderId()).Msg("Complete")
	return e.Paid(), nil
}

// Listen implements SmartContract
func (mockContract) Listen(ctx context.Context, out chan<- ContractSubmittedEvent) error {
	log.Ctx(ctx).Debug().Msg("Listening")
	defer log.Ctx(ctx).Debug().Msg("Stopping listening")

	sched := gocron.NewScheduler(time.UTC)
	_, err := sched.Every(1).Minute().Do(func() {
		e := &event{
			Order:       Order{ID: uuid.New()},
			attempts:    0,
			lastAttempt: time.Time{},
			state:       "Submitted",
			jobSpec: model.Spec{
				Engine:    model.EngineDocker,
				Verifier:  model.VerifierNoop,
				Publisher: model.PublisherEstuary,
				Docker: model.JobSpecDocker{
					Image:      "ubuntu",
					Entrypoint: []string{"date"},
				},
				Deal: model.Deal{
					Concurrency: 1,
				},
			},
		}
		log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("New order")
		out <- e
	})
	if err != nil {
		return err
	}

	sched.StartAsync()
	defer sched.Stop()

	<-ctx.Done()
	return nil
}

// Refund implements SmartContract
func (mockContract) Refund(ctx context.Context, e ContractSubmittedEvent) (ContractRefundedEvent, error) {
	log.Ctx(ctx).Debug().Stringer("id", e.OrderId()).Msg("Refunded")
	return e.Refunded(), nil
}

func MockContract() SmartContract {
	return mockContract{}
}
