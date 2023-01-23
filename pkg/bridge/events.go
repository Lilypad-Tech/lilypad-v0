package bridge

import (
	"time"

	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/google/uuid"
)

type Event interface {
	OrderId() uuid.UUID
}

type Retryable interface {
	Attempts() int
	AddAttempt() int
	LastAttempt() time.Time
}

type ContractSubmittedEvent interface {
	Event
	Retryable

	Spec() model.Spec

	Refunded() ContractRefundedEvent
	Paid() ContractPaidEvent
	JobCreated(*model.Job) BacalhauJobRunningEvent
}

type BacalhauJobRunningEvent interface {
	Event
	Retryable

	ContractSubmittedEvent

	JobID() string

	Completed() BacalhauJobCompletedEvent
	Failed() BacalhauJobFailedEvent
}

type BacalhauJobCompletedEvent interface {
	Event
	Retryable

	BacalhauJobRunningEvent
}

type BacalhauJobFailedEvent interface {
	Event
	Retryable

	BacalhauJobRunningEvent
}

type ContractPaidEvent interface {
	Event

	BacalhauJobCompletedEvent
}

type ContractRefundedEvent interface {
	Event

	ContractSubmittedEvent
}

type event struct {
	orderId     uuid.UUID
	attempts    int
	lastAttempt time.Time
	state       string
	jobSpec     model.Spec
	jobId       string
}

// The smart contract order ID.
func (e *event) OrderId() uuid.UUID {
	return e.orderId
}

// Log the event as being retried.
func (e *event) AddAttempt() int {
	e.attempts += 1
	e.lastAttempt = time.Now()
	return e.attempts
}

// Returns the number of times we have tried to process the event and move it
// into the next state.
func (e *event) Attempts() int {
	return e.attempts
}

// Returns the last time we tried to process the event.
func (e *event) LastAttempt() time.Time {
	return e.lastAttempt
}

// Records that a ContractSubmittedEvent has been sent to the Bacalhau network
// as a job.
func (e *event) JobCreated(job *model.Job) BacalhauJobRunningEvent {
	e.state = "JobCreated"
	e.jobId = job.Metadata.ID
	return e
}

// Records that a BacalhauJobCompletedEvent has been successfully sent to the
// smart contract for payment.
func (e *event) Paid() ContractPaidEvent {
	e.state = "Paid"
	return e
}

// Records that an Event has been successfully returned to the smart contract
// for a refund.
func (e *event) Refunded() ContractRefundedEvent {
	e.state = "Refunded"
	return e
}

// The Bacalhau job spec that the contract is asking us to run.
func (e *event) Spec() model.Spec {
	return e.jobSpec
}

// Records that a running Bacalhau job has completed.
func (e *event) Completed() BacalhauJobCompletedEvent {
	e.state = "JobCompleted"
	return e
}

// Records that a running Bacalhau job has failed.
func (e *event) Failed() BacalhauJobFailedEvent {
	return e
}

// The ID of the job on the Bacalhau network.
func (e *event) JobID() string {
	return e.jobId
}

var _ BacalhauJobRunningEvent = (*event)(nil)
var _ ContractSubmittedEvent = (*event)(nil)
