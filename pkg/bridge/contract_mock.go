package bridge

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

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
	spec := fastSpec
	spec.Docker.Entrypoint = append(spec.Docker.Entrypoint, "test")
	specJson, err := json.Marshal(spec)
	if err != nil {
		panic(err)
	}
	e := &event{
		attempts:    0,
		lastAttempt: time.Time{},
		state:       OrderStateSubmitted,
		jobSpec:     specJson,
	}
	id := make([]byte, 0, 32)
	_, err = rand.Read(id)
	if err != nil {
		panic(err)
	}
	copy(e.orderId[:], id[0:32])
	return e
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
