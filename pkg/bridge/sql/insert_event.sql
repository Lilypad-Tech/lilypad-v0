INSERT INTO events
	(orderId, attempts, lastAttempt, state, jobSpec, jobId)
    VALUES (:orderId, :attempts, :lastAttempt, :state, :jobSpec, :jobId);
