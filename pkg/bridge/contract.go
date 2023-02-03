package bridge

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"math/big"
	"time"

	"github.com/bacalhau-project/lilypad/hardhat/artifacts/contracts/LilypadEvents.sol"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

type SmartContract interface {
	Listen(context.Context, chan<- ContractSubmittedEvent) error

	Complete(context.Context, BacalhauJobCompletedEvent) (ContractPaidEvent, error)

	Refund(context.Context, ContractFailedEvent) (ContractRefundedEvent, error)
}

type realContract struct {
	client     *ethclient.Client
	contract   *LilypadEvents.LilypadEvents
	privateKey *ecdsa.PrivateKey

	maxSeenBlock uint64
}

func (r *realContract) publicKey() *ecdsa.PublicKey {
	return r.privateKey.Public().(*ecdsa.PublicKey)
}

func (r *realContract) wallet() common.Address {
	return crypto.PubkeyToAddress(*r.publicKey())
}

func (r *realContract) pendingNonce(ctx context.Context) (uint64, error) {
	return r.client.PendingNonceAt(ctx, r.wallet())
}

func (r *realContract) prepareTransaction(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := r.pendingNonce(ctx)
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(r.privateKey, big.NewInt(3141))
	if err != nil {
		return nil, err
	}

	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0)
	opts.Context = ctx

	return opts, nil
}

// Complete implements SmartContract
func (r *realContract) Complete(ctx context.Context, event BacalhauJobCompletedEvent) (ContractPaidEvent, error) {
	opts, err := r.prepareTransaction(ctx)
	if err != nil {
		return nil, err
	}

	txn, err := r.contract.LilypadEventsTransactor.ReturnBacalhauResults(
		opts,
		event.OrderRequestor(),
		event.OrderId().Big(),
		event.OrderName(),
		event.Result().String(),
	)
	if err != nil {
		return nil, err
	}

	log.Ctx(ctx).Debug().Stringer("txn", txn.Hash()).Msg("Results returned")
	return event.Paid(), nil
}

// Listen implements SmartContract
func (r *realContract) Listen(ctx context.Context, out chan<- ContractSubmittedEvent) error {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(15*time.Second).SingletonMode().Do(r.ReadLogs, ctx, out)
	if err != nil {
		return err
	}

	scheduler.StartAsync()
	defer scheduler.Stop()

	<-ctx.Done()
	return nil
}

func (r *realContract) ReadLogs(ctx context.Context, out chan<- ContractSubmittedEvent) {
	log.Ctx(ctx).Debug().Uint64("fromBlock", r.maxSeenBlock+1).Msg("Polling for smart contract events")

	opts := bind.FilterOpts{Start: uint64(r.maxSeenBlock + 1), Context: ctx}
	logs, err := r.contract.LilypadEventsFilterer.FilterNewBacalhauJobSubmitted(&opts, nil)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Send()
		return
	}
	defer logs.Close()

	for logs.Next() {
		recvEvent := logs.Event
		log.Ctx(ctx).Debug().
			Stringer("txn", recvEvent.Raw.TxHash).
			Uint64("block#", recvEvent.Raw.BlockNumber).
			Str("params", recvEvent.Params).
			Bool("removed", recvEvent.Raw.Removed).
			Msg("Event")

		if recvEvent.Raw.Removed {
			continue
		}

		specObj := fastSpec
		specObj.Docker.Entrypoint = append(specObj.Docker.Entrypoint, recvEvent.Params)
		spec, err := json.Marshal(specObj)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Send()
			continue
		}

		out <- &event{
			orderId:     recvEvent.Raw.TxHash.Bytes(),
			orderOwner:  recvEvent.RequestorContract.Bytes(),
			orderNumber: recvEvent.JobId.Int64(),
			orderName:   recvEvent.JobName,
			state:       OrderStateSubmitted,
			jobSpec:     spec,
		}

		r.maxSeenBlock = recvEvent.Raw.BlockNumber
	}
}

// Refund implements SmartContract
func (r *realContract) Refund(ctx context.Context, e ContractFailedEvent) (ContractRefundedEvent, error) {
	return e.Refunded(), nil
}

func NewContract(contractAddr common.Address, privateKey *ecdsa.PrivateKey) (SmartContract, error) {
	client, err := ethclient.Dial("wss://ws-filecoin-hyperspace.chainstacklabs.com/rpc/v0")
	if err != nil {
		return nil, err
	}

	contract, err := LilypadEvents.NewLilypadEvents(contractAddr, client)
	if err != nil {
		return nil, err
	}

	number, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	return &realContract{client, contract, privateKey, number}, nil
}

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
