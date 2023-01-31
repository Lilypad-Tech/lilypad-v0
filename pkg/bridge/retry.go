package bridge

import (
	"math"
	"time"
)

const defaultRetryTime = 5 * time.Second

var maxAttemptsByState map[OrderState]uint = map[OrderState]uint{
	OrderStateSubmitted: 5,
	OrderStateRunning:   0,
	OrderStateCompleted: 5,
	OrderStateJobError:  3,
	OrderStateFailed:    5,
	OrderStatePaid:      0,
	OrderStateRefunded:  0,
}

// A RetryStrategy defines how long an event should wait before it is retried.
type RetryStrategy func(e Retryable) time.Duration

// Exponential schedules the action to happen later into the future if more
// attempts have been made, with an exponential increase in the time waited.
var Exponential RetryStrategy = func(e Retryable) time.Duration {
	return time.Duration(math.Pow(defaultRetryTime.Seconds(), float64(e.Attempts()+1))) * time.Second
}

// Immediate schedules the action to happen again immediately, without waiting.
var Immediate RetryStrategy = func(e Retryable) time.Duration {
	return 0
}

func ShouldRetry(event Retryable) bool {
	return event.Attempts() < maxAttemptsByState[event.OrderState()]
}
