CREATE TABLE IF NOT EXISTS events (
    eventId     INTEGER PRIMARY KEY AUTOINCREMENT,
	orderId     VARCHAR(16),
	attempts    SMALLINT,
	lastAttempt VARCHAR(25),
	state       SMALLINT,
	jobSpec     BLOB,
	jobId       TEXT
);

CREATE UNIQUE INDEX IF NOT EXISTS event_versions ON events (orderId, eventId);

CREATE VIEW IF NOT EXISTS latest_events (eventId, orderId, attempts, lastAttempt, state, jobSpec, jobId) AS
    WITH events_with_max AS (
        SELECT *, LAST_VALUE(eventId) OVER (PARTITION BY orderId ORDER BY eventId RANGE BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING) AS maxEventId FROM events
    )
    SELECT eventId, orderId, attempts, lastAttempt, state, jobSpec, jobId
    FROM events_with_max
    WHERE eventId = maxEventId;
