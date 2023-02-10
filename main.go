package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/bacalhau-project/lilypad/pkg/bridge"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := log.Logger.WithContext(context.Background())

	lvl := zerolog.InfoLevel
	if level, found := os.LookupEnv("LOG_LEVEL"); found {
		var err error
		lvl, err = zerolog.ParseLevel(level)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	}
	zerolog.SetGlobalLevel(lvl)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	repo, err := bridge.NewSQLiteRepository(ctx, "lilypad.sqlite")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	addr := common.HexToAddress(os.Getenv("DEPLOYED_CONTRACT_ADDRESS"))
	privKey, err := crypto.HexToECDSA(os.Getenv("WALLET_PRIVATE_KEY"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "WALLET_PRIVATE_KEY: "+err.Error())
		return
	}

	contract, err := bridge.NewContract(addr, privKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	workflow := bridge.NewWorkflow(bridge.NewJobRunner(), contract, repo)

	err = workflow.Start(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
