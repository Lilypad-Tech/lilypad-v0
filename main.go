package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/bacalhau-project/lilypad/pkg/bridge"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := log.Logger.WithContext(context.Background())
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	immediateStop, immediateCancel := signal.NotifyContext(ctx, os.Kill)
	gracefulStop, gracefulCancel := signal.NotifyContext(immediateStop, os.Interrupt)
	defer gracefulCancel()
	defer immediateCancel()

	workflow := bridge.Workflow{
		Bacalhau: bridge.NewJobRunner(),
		Contract: bridge.TimerContract(),
	}

	err := workflow.Run(immediateStop, gracefulStop, immediateCancel)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
