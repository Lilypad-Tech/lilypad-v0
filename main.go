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

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	repo, err := bridge.NewSQLiteRepository(ctx, "lilypad.sqlite")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	workflow := bridge.NewWorkflow(bridge.NewJobRunner(), bridge.TimerContract(), repo)

	err = workflow.Start(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
