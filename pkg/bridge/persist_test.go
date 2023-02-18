package bridge

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
)

func repository(t *testing.T) Repository {
	repo, err := NewSQLiteRepository(context.Background(), filepath.Join(t.TempDir(), "test.sqlite"))
	require.NoError(t, err)
	return repo
}

func TestCanSaveAndReloadInEveryState(t *testing.T) {
	runTest := func(e Event, state OrderState) {
		repo := repository(t)
		require.NoError(t, repo.Save(e))

		for _, orderState := range OrderStates() {
			events, err := repo.Reload(orderState)
			require.NoError(t, err)
			if state == orderState {
				require.Len(t, events, 1)
				require.Equal(t, events[0].OrderId(), e.OrderId())
				require.Equal(t, events[0].OrderState(), e.OrderState())
			} else {
				require.Len(t, events, 0)
			}
		}

	}

	t.Run("ContractSubmittedEvent", func(t *testing.T) {
		runTest(exampleEvent(), OrderStateSubmitted)
	})

	t.Run("BacalhauJobRunningEvent", func(t *testing.T) {
		runTest(exampleEvent().JobCreated(model.NewJob()), OrderStateRunning)
	})

	t.Run("BacalhauJobCompletedEvent", func(t *testing.T) {
		runTest(exampleEvent().JobCreated(model.NewJob()).Completed(cid.Cid{}, "", "", 0), OrderStateCompleted)
	})

	t.Run("BacalhauJobFailedEvent", func(t *testing.T) {
		runTest(exampleEvent().JobCreated(model.NewJob()).JobError(""), OrderStateJobError)
	})

	t.Run("ContractPaidEvent", func(t *testing.T) {
		runTest(exampleEvent().JobCreated(model.NewJob()).Completed(cid.Cid{}, "", "", 0).Paid(), OrderStatePaid)
	})

	t.Run("ContractRefundedEvent", func(t *testing.T) {
		runTest(exampleEvent().JobCreated(model.NewJob()).Failed("").Refunded(), OrderStateRefunded)
	})
}

func TestOldVersionsAreNotReloaded(t *testing.T) {
	repo := repository(t)
	e := exampleEvent()
	require.NoError(t, repo.Save(e))

	e = e.JobCreated(model.NewJob())
	require.NoError(t, repo.Save(e))

	events, err := repo.Reload(OrderStateSubmitted)
	require.NoError(t, err)
	require.Empty(t, events)

	e = e.JobCreated(model.NewJob()).Completed(cid.Cid{}, "", "", 0).Paid()
	require.NoError(t, repo.Save(e))

	events, err = repo.Reload(OrderStateRunning)
	require.NoError(t, err)
	require.Empty(t, events)

	e = e.JobCreated(model.NewJob()).Failed("").Refunded()
	require.NoError(t, repo.Save(e))

	events, err = repo.Reload(OrderStatePaid)
	require.NoError(t, err)
	require.Empty(t, events)
}
