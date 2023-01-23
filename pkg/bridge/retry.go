package bridge

import (
	"context"
	"math"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

const defaultRetryTime = 5 * time.Second

type RetryableEvent interface {
	Event
	Retryable
}

func Retry[E RetryableEvent](ctx context.Context, maxAttempts int, in <-chan E, retry chan<- E, cancel chan<- E) {
	log.Ctx(ctx).Debug().Msg("Started")
	defer log.Ctx(ctx).Debug().Msg("Stopped")

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.StartAsync()
	defer scheduler.Stop()

	for {
		select {
		case e := <-in:
			ctx := log.Ctx(ctx).With().Stringer("id", e.OrderId()).Logger().WithContext(ctx)

			if e.Attempts() >= maxAttempts {
				cancel <- e
			} else {
				waitTime := time.Duration(math.Pow(defaultRetryTime.Seconds(), float64(e.Attempts()))) * time.Second
				nextTime := time.Now().Add(waitTime)
				log.Ctx(ctx).Debug().Dur("retry", waitTime).Msg("Scheduling retry")

				_, err := scheduler.WaitForSchedule().Every(1).StartAt(nextTime).LimitRunsTo(1).Do(func() {
					e.AddAttempt()
					retry <- e
				})
				if err != nil {
					log.Ctx(ctx).Error().Err(err).Msg("error scheduling retry")
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
