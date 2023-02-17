package bridge

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func Query(name string) string {
	sql, err := sqlFiles.ReadFile(fmt.Sprintf("sql/%s.sql", name))
	if err != nil {
		panic(err)
	}
	return string(sql)
}

type Repository interface {
	Save(Event) error

	Reload(OrderState) ([]Event, error)

	Exists(Event) (bool, error)
}

type sqlRepository struct {
	db *sql.DB

	insertEvent    *sql.Stmt
	eventExists    *sql.Stmt
	retrieveEvents *sql.Stmt
}

// Reload implements Repository
func (repo *sqlRepository) Reload(state OrderState) ([]Event, error) {
	rows, err := repo.retrieveEvents.Query(sql.Named("state", state))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]Event, 0)
	for rows.Next() {
		var e event
		var lastAttemptString string
		err = rows.Scan(
			&e.eventId,
			&e.orderId,
			&e.orderOwner,
			&e.orderNumber,
			&e.orderResultType,
			&e.attempts,
			&lastAttemptString,
			&e.state,
			&e.jobSpec,
			&e.jobId,
			&e.jobResult,
			&e.jobStdout,
			&e.jobStderr,
			&e.jobExitcode,
		)
		if err != nil {
			break
		}
		e.lastAttempt, err = time.Parse(time.RFC3339, lastAttemptString)
		if err != nil {
			break
		}
		events = append(events, &e)
	}

	return events, err
}

// Save implements Repository
func (repo *sqlRepository) Save(in Event) error {
	e, ok := in.(*event) // TODO smelly
	if !ok {
		return fmt.Errorf("don't know how to save event of type %T", in)
	}
	_, err := repo.insertEvent.Exec(
		sql.Named("orderId", e.orderId),
		sql.Named("orderOwner", e.orderOwner),
		sql.Named("orderNumber", e.orderNumber),
		sql.Named("orderResultType", e.orderResultType),
		sql.Named("attempts", e.attempts),
		sql.Named("lastAttempt", e.lastAttempt.Format(time.RFC3339)),
		sql.Named("state", e.state),
		sql.Named("jobSpec", e.jobSpec),
		sql.Named("jobId", e.jobId),
		sql.Named("jobResult", e.jobResult),
		sql.Named("jobStdout", e.jobStdout),
		sql.Named("jobStderr", e.jobStderr),
		sql.Named("jobExitcode", e.jobExitcode),
	)
	return err
}

func (repo *sqlRepository) Exists(in Event) (bool, error) {
	res, err := repo.eventExists.Query(sql.Named("orderId", in.OrderId()))
	if err != nil {
		return false, err
	}
	return res.Next(), nil
}

func NewSQLiteRepository(ctx context.Context, path string) (Repository, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(Query("init"))
	if err != nil {
		return nil, err
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	insertEvent, err := conn.PrepareContext(ctx, Query("insert_event"))
	if err != nil {
		return nil, err
	}

	eventExists, err := conn.PrepareContext(ctx, Query("event_exists"))
	if err != nil {
		return nil, err
	}

	retrieveEvents, err := conn.PrepareContext(ctx, Query("retrieve_events"))
	if err != nil {
		return nil, err
	}

	return &sqlRepository{
		db:             db,
		insertEvent:    insertEvent,
		eventExists:    eventExists,
		retrieveEvents: retrieveEvents,
	}, nil
}

func Reload[E Event](repo Repository, state OrderState) ([]E, error) {
	events, err := repo.Reload(state)
	if err != nil {
		return nil, err
	}

	casts := make([]E, 0, len(events))
	for _, event := range events {
		casts = append(casts, event.(E))
	}
	return casts, nil
}

func ReloadToChan[E Event](repo Repository, state OrderState, out chan<- Event) error {
	events, err := Reload[E](repo, state)
	if err != nil {
		return err
	}
	for _, event := range events {
		out <- event
	}
	return nil
}
