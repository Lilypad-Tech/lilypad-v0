package bridge

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestExponentialStrategy(t *testing.T) {
	testCases := []time.Duration{
		5 * time.Second,
		25 * time.Second,
		125 * time.Second,
		625 * time.Second,
	}

	e := exampleEvent()
	for attempts, wait := range testCases {
		require.Equal(t, wait, Exponential(e),
			fmt.Sprintf("%d attempts should wait %s", attempts, wait))
		e.AddAttempt()
	}
}

func TestImmediateStrategy(t *testing.T) {
	e := exampleEvent()
	for i := 0; i <= 3; i++ {
		require.Equal(t, 0*time.Second, Immediate(e),
			fmt.Sprintf("%d attempts should fire immediately", i))
	}
}

func TestShouldRetry(t *testing.T) {
	e := exampleEvent()
	for i := 0; i <= 6; i++ {
		require.Equal(t, i < 5, ShouldRetry(e), i)
		e.AddAttempt()
	}
}
