CREATE TABLE IF NOT EXISTS events (
    eventId     INTEGER PRIMARY KEY AUTOINCREMENT,
	orderId     VARCHAR(32),
	orderOwner  VARCHAR(32),
	orderNumber BIGINT,
	orderResultType INTEGER,
	attempts    SMALLINT,
	lastAttempt VARCHAR(25),
	state       SMALLINT,
	jobSpec     BLOB,
	jobId       TEXT,
	jobResult   TEXT,
	jobStdout	TEXT,
	jobStderr	TEXT,
	jobExitcode SMALLINT
);

CREATE UNIQUE INDEX IF NOT EXISTS event_versions ON events (orderId, eventId);

CREATE VIEW IF NOT EXISTS latest_events AS
    WITH events_with_max AS (
        SELECT *, LAST_VALUE(eventId) OVER (PARTITION BY orderId ORDER BY eventId RANGE BETWEEN CURRENT ROW AND UNBOUNDED FOLLOWING) AS maxEventId FROM events
    )
    SELECT *
    FROM events_with_max
    WHERE eventId = maxEventId;
