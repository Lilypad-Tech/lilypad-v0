package bridge

import (
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/ipfs/go-cid"
)

//go:generate stringer -type=OrderState --trimprefix=OrderState
type OrderState int

const (
	OrderStateSubmitted OrderState = iota
	OrderStateRunning
	OrderStateCompleted
	OrderStatePaid
	OrderStateRefunded
	OrderStateJobError
	OrderStateFailed
)

func OrderStates() [7]OrderState {
	return [7]OrderState{
		OrderStateSubmitted,
		OrderStateRunning,
		OrderStateCompleted,
		OrderStatePaid,
		OrderStateRefunded,
		OrderStateJobError,
		OrderStateFailed,
	}
}

type ResultType uint8

const (
	ResultTypeCID      ResultType = 0
	ResultTypeStdOut   ResultType = 1
	ResultTypeStdErr   ResultType = 2
	ResultTypeExitCode ResultType = 3
)

func (r ResultType) Valid() bool {
	return r >= ResultTypeCID && r <= ResultTypeExitCode
}

type Event interface {
	OrderId() common.Hash
	OrderState() OrderState
}

type Retryable interface {
	Event

	Attempts() uint
	AddAttempt() uint
	LastAttempt() time.Time
}

type ContractSubmittedEvent interface {
	Event
	Retryable

	OrderNumber() int64
	OrderResultType() ResultType
	OrderRequestor() common.Address
	Spec() (model.Spec, error)

	Failed(err string) ContractFailedEvent
	JobCreated(*model.Job) BacalhauJobRunningEvent
}

type BacalhauJobRunningEvent interface {
	Event
	Retryable

	ContractSubmittedEvent

	JobID() string

	Completed(result cid.Cid, stdout, stderr string, exitcode int) BacalhauJobCompletedEvent
	JobError(err string) BacalhauJobFailedEvent
}

type BacalhauJobCompletedEvent interface {
	Event
	Retryable

	BacalhauJobRunningEvent

	Result() cid.Cid
	StdOut() string
	StdErr() string
	ExitCode() int

	Paid() ContractPaidEvent
}

type BacalhauJobFailedEvent interface {
	Event
	Retryable

	BacalhauJobRunningEvent

	Error() string

	Retry() ContractSubmittedEvent
}

type ContractFailedEvent interface {
	Event
	Retryable

	ContractSubmittedEvent

	Error() string

	Refunded() ContractRefundedEvent
}

type ContractPaidEvent interface {
	Event

	BacalhauJobCompletedEvent
}

type ContractRefundedEvent interface {
	Event

	ContractFailedEvent
}

type event struct {
	eventId         uint
	orderId         []byte
	orderOwner      []byte
	orderNumber     int64
	orderResultType uint8
	attempts        uint
	lastAttempt     time.Time
	state           OrderState
	jobSpec         []byte
	jobId           string
	jobResult       string
	jobStdout       string
	jobStderr       string
	jobExitcode     int
}

// The smart contract order ID.
func (e *event) OrderId() common.Hash {
	return common.BytesToHash(e.orderId)
}

func (e *event) OrderState() OrderState {
	return e.state
}

// OrderName implements BacalhauJobCompletedEvent
func (e *event) OrderResultType() ResultType {
	return ResultType(e.orderResultType)
}

// OrderNumber implements BacalhauJobCompletedEvent
func (e *event) OrderNumber() int64 {
	return e.orderNumber
}

func (e *event) OrderRequestor() common.Address {
	return common.BytesToAddress(e.orderOwner)
}

// Result implements BacalhauJobCompletedEvent
func (e *event) Result() cid.Cid {
	return cid.MustParse(e.jobResult)
}

// Log the event as being retried.
func (e *event) AddAttempt() uint {
	e.attempts += 1
	e.lastAttempt = time.Now()
	return e.attempts
}

// Returns the number of times we have tried to process the event and move it
// into the next state.
func (e *event) Attempts() uint {
	return e.attempts
}

// Returns the last time we tried to process the event.
func (e *event) LastAttempt() time.Time {
	return e.lastAttempt
}

// ExitCode implements BacalhauJobCompletedEvent
func (e *event) ExitCode() int {
	return e.jobExitcode
}

// StdErr implements BacalhauJobCompletedEvent
func (e *event) StdErr() string {
	return e.jobStderr
}

// StdOut implements BacalhauJobCompletedEvent
func (e *event) StdOut() string {
	return e.jobStdout
}

// Error implements BacalhauJobFailedEvent
func (e *event) Error() string {
	return e.jobStderr
}

// Records that a ContractSubmittedEvent has been sent to the Bacalhau network
// as a job.
func (e *event) JobCreated(job *model.Job) BacalhauJobRunningEvent {
	e.state = OrderStateRunning
	e.jobId = job.Metadata.ID
	return e
}

// Records that a BacalhauJobCompletedEvent has been successfully sent to the
// smart contract for payment.
func (e *event) Paid() ContractPaidEvent {
	e.state = OrderStatePaid
	return e
}

// Records that an Event has been successfully returned to the smart contract
// for a refund.
func (e *event) Refunded() ContractRefundedEvent {
	e.state = OrderStateRefunded
	return e
}

// The Bacalhau job spec that the contract is asking us to run.
func (e *event) Spec() (spec model.Spec, err error) {
	err = json.Unmarshal(e.jobSpec, &spec)
	return
}

// Records that a running Bacalhau job has completed.
func (e *event) Completed(result cid.Cid, stdout, stderr string, exitcode int) BacalhauJobCompletedEvent {
	e.state = OrderStateCompleted
	e.jobResult = result.String()
	e.jobStdout = stdout
	e.jobStderr = stderr
	e.jobExitcode = exitcode
	return e
}

// Records that a running Bacalhau job has failed.
func (e *event) JobError(err string) BacalhauJobFailedEvent {
	e.state = OrderStateJobError
	e.jobStderr = err
	return e
}

// Records that an errored Bacalhau job is being retried.
func (e *event) Retry() ContractSubmittedEvent {
	e.state = OrderStateSubmitted
	e.AddAttempt()
	return e
}

// Records that a contract has failed permanently.
func (e *event) Failed(err string) ContractFailedEvent {
	e.state = OrderStateFailed
	e.jobStderr = err
	return e
}

// The ID of the job on the Bacalhau network.
func (e *event) JobID() string {
	return e.jobId
}

var _ BacalhauJobRunningEvent = (*event)(nil)
var _ ContractSubmittedEvent = (*event)(nil)
