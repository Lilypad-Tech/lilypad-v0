SELECT eventId, orderId, attempts, lastAttempt, state, jobSpec, jobId
FROM latest_events
WHERE state = :state;
