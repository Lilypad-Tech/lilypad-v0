package bridge

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

type WorkflowTestSuite struct {
	suite.Suite

	completed chan ContractPaidEvent
	refunded  chan ContractRefundedEvent

	selectCancel context.CancelFunc

	workflowGroup  errgroup.Group
	workflowCancel context.CancelFunc
}

func TestWorkflowSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

func (suite *WorkflowTestSuite) SetupSuite() {
	suite.NoError(system.InitConfigForTesting(suite.T()))
	GlobalRetryStrategy = Immediate
}

func (suite *WorkflowTestSuite) SetupTest() {
	suite.completed = make(chan ContractPaidEvent, 1)
	suite.refunded = make(chan ContractRefundedEvent, 1)
}

func (suite *WorkflowTestSuite) RunWorkflow(w *Workflow) {
	ctx, engineCancel := context.WithCancel(context.Background())
	suite.workflowCancel = engineCancel

	suite.workflowGroup.Go(func() error {
		return w.Run(ctx, ctx, engineCancel)
	})
}

func (suite *WorkflowTestSuite) Timeout() <-chan struct{} {
	timeout, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	suite.selectCancel = cancel
	return timeout.Done()
}

func (suite *WorkflowTestSuite) TearDownTest() {
	suite.workflowCancel()
	suite.NoError(suite.workflowGroup.Wait())
}

func (suite *WorkflowTestSuite) SuccessfulComplete() ContractCompleteHandler {
	return func(ctx context.Context, bjce BacalhauJobCompletedEvent) (ContractPaidEvent, error) {
		suite.completed <- bjce.Paid()
		return bjce.Paid(), nil
	}
}

func (suite *WorkflowTestSuite) ErrorComplete() ContractCompleteHandler {
	return func(ctx context.Context, bjce BacalhauJobCompletedEvent) (ContractPaidEvent, error) {
		return nil, errors.New("failed to pay the smart contract")
	}
}

func (suite *WorkflowTestSuite) SuccessfulRefund() ContractRefundHandler {
	return func(ctx context.Context, cse ContractSubmittedEvent) (ContractRefundedEvent, error) {
		suite.refunded <- cse.Refunded()
		return cse.Refunded(), nil
	}
}

func (suite *WorkflowTestSuite) EmitOne(e ContractSubmittedEvent) ContractListenHandler {
	return func(ctx context.Context, c chan<- ContractSubmittedEvent) error {
		c <- e
		return nil
	}
}

func (suite *WorkflowTestSuite) TestHappyPath() {
	e := exampleEvent()

	suite.RunWorkflow(&Workflow{
		Bacalhau: &mockRunner{
			CreateHandler:        SuccessfulCreate,
			FindCompletedHandler: SuccssfulFind,
		},
		Contract: &mockContract{
			CompleteHandler: suite.SuccessfulComplete(),
			RefundHandler:   suite.SuccessfulRefund(),
			ListenHandler:   suite.EmitOne(e),
		},
	})

	select {
	case result := <-suite.completed:
		suite.Equal(e.OrderId(), result.OrderId())
	case <-suite.refunded:
		suite.Fail("Should not have got a refunded event")
	case <-suite.Timeout():
		suite.Fail("Timed out")
	}
}

func (suite *WorkflowTestSuite) TestRefundsOnFail() {
	e := exampleEvent()

	runner := &mockRunner{
		CreateHandler:        SuccessfulCreate,
		FindCompletedHandler: SuccssfulFind,
	}
	contract := &mockContract{
		CompleteHandler: suite.SuccessfulComplete(),
		RefundHandler:   suite.SuccessfulRefund(),
		ListenHandler:   suite.EmitOne(e),
	}
	w := &Workflow{Bacalhau: runner, Contract: contract}

	runTest := func() {
		suite.RunWorkflow(w)

		select {
		case <-suite.completed:
			suite.Fail("Should not have got a completed event")
		case result := <-suite.refunded:
			suite.Equal(e.OrderId(), result.OrderId())
		case <-suite.Timeout():
			suite.Fail("Timed out")
		}
	}

	suite.Run("Create", func() {
		runner.CreateHandler = ErrorCreate
		runTest()
		runner.CreateHandler = SuccessfulCreate
		suite.TearDownTest()
	})

	suite.Run("FindCompleted", func() {
		runner.FindCompletedHandler = FailedFind
		runTest()
		runner.FindCompletedHandler = SuccssfulFind
		suite.TearDownTest()
	})

	suite.Run("Paid", func() {
		contract.CompleteHandler = suite.ErrorComplete()
		runTest()
		contract.CompleteHandler = suite.SuccessfulComplete()
		suite.TearDownTest()
	})
}
