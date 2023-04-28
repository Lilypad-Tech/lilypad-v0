package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/bacalhau-project/bacalhau/pkg/logger"
	"github.com/bacalhau-project/lilypad/pkg/bridge"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func EnvOrDefault(env, defaultValue string) (value string) {
	value, found := os.LookupEnv(env)
	if !found {
		value = defaultValue
	}
	return
}

func main() {
	logType, err := logger.ParseLogMode(EnvOrDefault("LOG_MODE", "default"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	logger.ConfigureLogging(logType)

	lvl, err := zerolog.ParseLevel(EnvOrDefault("LOG_LEVEL", "INFO"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	zerolog.SetGlobalLevel(lvl)

	ctx := log.Logger.WithContext(context.Background())
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	sqliteFileLocation := EnvOrDefault("SQLITE_FILE_LOCATION", "lilypad.sqlite")
	repo, err := bridge.NewSQLiteRepository(ctx, sqliteFileLocation)
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
