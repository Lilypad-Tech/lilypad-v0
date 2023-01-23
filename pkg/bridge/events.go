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
	Order       Order
	attempts    int
	lastAttempt time.Time
	state       string
	jobSpec     model.Spec
	jobId       string
}

// OrderId implements BacalhauJobRunningEvent
func (e *event) OrderId() uuid.UUID {
	return e.Order.ID
}

// AddAttempt implements BacalhauJobRunningEvent
func (e *event) AddAttempt() int {
	e.attempts += 1
	e.lastAttempt = time.Now()
	return e.attempts
}

// Attempts implements BacalhauJobRunningEvent
func (e *event) Attempts() int {
	return e.attempts
}

// LastAttempt implements BacalhauJobRunningEvent
func (e *event) LastAttempt() time.Time {
	return e.lastAttempt
}

// JobCreated implements BacalhauJobRunningEvent
func (e *event) JobCreated(job *model.Job) BacalhauJobRunningEvent {
	e.state = "JobCreated"
	e.jobId = job.Metadata.ID
	return e
}

// Paid implements BacalhauJobRunningEvent
func (e *event) Paid() ContractPaidEvent {
	e.state = "Paid"
	return e
}

// Refunded implements BacalhauJobRunningEvent
func (e *event) Refunded() ContractRefundedEvent {
	e.state = "Refunded"
	return e
}

// Spec implements BacalhauJobRunningEvent
func (e *event) Spec() model.Spec {
	return e.jobSpec
}

// Completed implements BacalhauJobRunningEvent
func (e *event) Completed() BacalhauJobCompletedEvent {
	e.state = "JobCompleted"
	return e
}

// Failed implements BacalhauJobRunningEvent
func (e *event) Failed() BacalhauJobFailedEvent {
	return e
}

// JobID implements BacalhauJobRunningEvent
func (e *event) JobID() string {
	return e.jobId
}

var _ BacalhauJobRunningEvent = (*event)(nil)
var _ ContractSubmittedEvent = (*event)(nil)
