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
)

func main() {
	// bacalhaus sets the default logger using buffered writer, so here just create a new one to ensure logs printed.
	defaultLog := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx := defaultLog.WithContext(context.Background())

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

	if !common.IsHexAddress(os.Getenv("DEPLOYED_CONTRACT_ADDRESS")) {
		fmt.Fprintln(os.Stderr, "DEPLOYED_CONTRACT_ADDRESS: invalid hex address string")
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
