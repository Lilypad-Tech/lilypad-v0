SELECT eventId, orderId, orderOwner, orderNumber, orderResultType, attempts, lastAttempt, state, jobSpec, jobId, jobResult, jobStdout, jobStderr, jobExitcode
FROM latest_events
WHERE state = :state;
