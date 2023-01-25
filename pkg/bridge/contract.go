package bridge

import (
	"context"
	"encoding/json"
	"time"

	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type SmartContract interface {
	Listen(context.Context, chan<- ContractSubmittedEvent) error

	Complete(context.Context, BacalhauJobCompletedEvent) (ContractPaidEvent, error)

	Refund(context.Context, ContractFailedEvent) (ContractRefundedEvent, error)
}

type ContractCompleteHandler func(context.Context, BacalhauJobCompletedEvent) (ContractPaidEvent, error)
type ContractRefundHandler func(context.Context, ContractFailedEvent) (ContractRefundedEvent, error)
type ContractListenHandler func(ctx context.Context, c chan<- ContractSubmittedEvent) error

type mockContract struct {
	CompleteHandler ContractCompleteHandler
	RefundHandler   ContractRefundHandler
	ListenHandler   ContractListenHandler
}

// Complete implements SmartContract
func (mock mockContract) Complete(ctx context.Context, e BacalhauJobCompletedEvent) (ContractPaidEvent, error) {
	log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("Complete")
	if mock.CompleteHandler != nil {
		return mock.CompleteHandler(ctx, e)
	}
	return e.Paid(), nil
}

// Listen implements SmartContract
func (mock mockContract) Listen(ctx context.Context, out chan<- ContractSubmittedEvent) error {
	log.Ctx(ctx).Debug().Msg("Listening")
	defer log.Ctx(ctx).Debug().Msg("Stopping listening")

	if mock.ListenHandler != nil {
		return mock.ListenHandler(ctx, out)
	}
	return nil
}

func exampleEvent() ContractSubmittedEvent {
	spec, err := json.Marshal(model.Spec{
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
	})
	if err != nil {
		panic(err)
	}
	return &event{
		orderId:     uuid.New(),
		attempts:    0,
		lastAttempt: time.Time{},
		state:       OrderStateSubmitted,
		jobSpec:     spec,
	}
}

// Refund implements SmartContract
func (mock mockContract) Refund(ctx context.Context, e ContractFailedEvent) (ContractRefundedEvent, error) {
	log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("Refunded")
	if mock.RefundHandler != nil {
		return mock.RefundHandler(ctx, e)
	}
	return e.Refunded(), nil
}

func TimerContract() SmartContract {
	return mockContract{
		ListenHandler: func(ctx context.Context, out chan<- ContractSubmittedEvent) error {
			sched := gocron.NewScheduler(time.UTC)
			_, err := sched.WaitForSchedule().Every(10).Second().Do(func() {
				e := exampleEvent()
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
		},
	}
}
