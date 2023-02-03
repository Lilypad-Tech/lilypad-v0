SELECT eventId, orderId, orderOwner, orderNumber, orderName, attempts, lastAttempt, state, jobSpec, jobId, jobResult
FROM latest_events
WHERE state = :state;
