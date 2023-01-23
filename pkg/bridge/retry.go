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

type RetryStrategy func(e RetryableEvent, scheduler *gocron.Scheduler) *gocron.Scheduler

var Exponential RetryStrategy = func(e RetryableEvent, scheduler *gocron.Scheduler) *gocron.Scheduler {
	waitTime := time.Duration(math.Pow(defaultRetryTime.Seconds(), float64(e.Attempts()))) * time.Second
	nextTime := time.Now().Add(waitTime)

	return scheduler.WaitForSchedule().Every(1).StartAt(nextTime)
}

var Immediate RetryStrategy = func(e RetryableEvent, scheduler *gocron.Scheduler) *gocron.Scheduler {
	return scheduler.Every(1).StartImmediately()
}

var GlobalRetryStrategy = Exponential

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
				schedule := GlobalRetryStrategy(e, scheduler)
				_, err := schedule.LimitRunsTo(1).Do(func() {
					e.AddAttempt()
					retry <- e
				})

				_, nextRun := schedule.NextRun()
				log.Ctx(ctx).Debug().Time("at", nextRun).Dur("wait", time.Until(nextRun)).Msg("Scheduling retry")

				if err != nil {
					log.Ctx(ctx).Error().Err(err).Msg("error scheduling retry")
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
