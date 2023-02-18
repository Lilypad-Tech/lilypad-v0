package bridge

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

type WorkflowTestSuite struct {
	suite.Suite

	completed chan ContractPaidEvent
	refunded  chan ContractRefundedEvent

	selectCancel context.CancelFunc

	workflowCtx    context.Context
	workflowGroup  errgroup.Group
	workflowCancel context.CancelFunc
}

func TestWorkflowSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

func (suite *WorkflowTestSuite) SetupSuite() {
	suite.NoError(system.InitConfigForTesting(suite.T()))
	defaultRetryStrategy = Immediate
	defaultJobCheckInterval = 20 * time.Millisecond
}

func (suite *WorkflowTestSuite) SetupTest() {
	ctx, cancel := context.WithCancel(context.Background())
	suite.workflowCtx = ctx
	suite.workflowCancel = cancel
	suite.completed = make(chan ContractPaidEvent, 1)
	suite.refunded = make(chan ContractRefundedEvent, 1)
}

func (suite *WorkflowTestSuite) RunWorkflow(w *Workflow) {
	suite.workflowGroup.Go(func() error {
		return w.Start(suite.workflowCtx)
	})
}

func (suite *WorkflowTestSuite) Timeout() <-chan struct{} {
	timeout, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	suite.selectCancel = cancel
	return timeout.Done()
}

func (suite *WorkflowTestSuite) Repository() Repository {
	tempDir := suite.T().TempDir()
	repo, err := NewSQLiteRepository(suite.workflowCtx, filepath.Join(tempDir, "test.sqlite"))
	suite.NoError(err)
	return repo
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
	return func(ctx context.Context, cse ContractFailedEvent) (ContractRefundedEvent, error) {
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

func (suite *WorkflowTestSuite) EmitNone() ContractListenHandler {
	return func(ctx context.Context, c chan<- ContractSubmittedEvent) error {
		return nil
	}
}

func (suite *WorkflowTestSuite) TestHappyPath() {
	e := exampleEvent()

	suite.RunWorkflow(NewWorkflow(
		&mockRunner{
			CreateHandler:        SuccessfulCreate,
			FindCompletedHandler: SuccssfulFind,
		},
		&mockContract{
			CompleteHandler: suite.SuccessfulComplete(),
			RefundHandler:   suite.SuccessfulRefund(),
			ListenHandler:   suite.EmitOne(e),
		},
		suite.Repository(),
	))

	select {
	case result := <-suite.completed:
		suite.Equal(e.OrderId(), result.OrderId())
	case <-suite.refunded:
		suite.Fail("Should not have got a refunded event")
	case <-suite.Timeout():
		suite.Fail("Timed out")
	}
}

func (suite *WorkflowTestSuite) RefundOnFailTest(
	create RunnerCreateHandler,
	find RunnerFindCompletedHandler,
	complete ContractCompleteHandler,
) {
	e := exampleEvent()
	w := NewWorkflow(
		&mockRunner{
			CreateHandler:        create,
			FindCompletedHandler: find,
		},
		&mockContract{
			CompleteHandler: complete,
			RefundHandler:   suite.SuccessfulRefund(),
			ListenHandler:   suite.EmitOne(e),
		},
		suite.Repository(),
	)

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

func (suite *WorkflowTestSuite) TestCreateRefunded() {
	suite.RefundOnFailTest(ErrorCreate, SuccssfulFind, suite.SuccessfulComplete())
}

func (suite *WorkflowTestSuite) TestFindCompletedRefunded() {
	suite.RefundOnFailTest(SuccessfulCreate, FailedFind, suite.SuccessfulComplete())
}

func (suite *WorkflowTestSuite) TestPaidRefunded() {
	suite.RefundOnFailTest(SuccessfulCreate, SuccssfulFind, suite.ErrorComplete())
}

func (suite *WorkflowTestSuite) ReloadEventTest(event Event) {
	repo := suite.Repository()
	suite.NoError(repo.Save(event))

	suite.RunWorkflow(NewWorkflow(
		&mockRunner{
			CreateHandler:        SuccessfulCreate,
			FindCompletedHandler: SuccssfulFind,
		},
		&mockContract{
			CompleteHandler: suite.SuccessfulComplete(),
			RefundHandler:   suite.SuccessfulRefund(),
			ListenHandler:   suite.EmitNone(),
		},
		repo,
	))

	select {
	case e := <-suite.completed:
		suite.Equal(event.OrderId(), e.OrderId())
	case e := <-suite.refunded:
		suite.Equal(event.OrderId(), e.OrderId())
	case <-suite.Timeout():
		suite.Fail("Timed out")
	}
}

func (suite *WorkflowTestSuite) TestSubmittedEventsAreReloaded() {
	suite.ReloadEventTest(exampleEvent())
}

func (suite *WorkflowTestSuite) TestRunningEventsAreReloaded() {
	suite.ReloadEventTest(exampleEvent().JobCreated(model.NewJob()))
}

func (suite *WorkflowTestSuite) TestCompletedEventsAreReloaded() {
	suite.ReloadEventTest(exampleEvent().JobCreated(model.NewJob()).Completed(cid.Cid{}, "", "", 0))
}

func (suite *WorkflowTestSuite) TestErroredEventsAreReloaded() {
	suite.ReloadEventTest(exampleEvent().JobCreated(model.NewJob()).JobError(""))
}

func (suite *WorkflowTestSuite) TestFailedEventsAreReloaded() {
	suite.ReloadEventTest(exampleEvent().JobCreated(model.NewJob()).Failed(""))
}
