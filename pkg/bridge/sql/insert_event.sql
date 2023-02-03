INSERT INTO events
	(orderId, orderOwner, orderNumber, orderName, attempts, lastAttempt, state, jobSpec, jobId, jobResult)
    VALUES (:orderId, :orderOwner, :orderNumber, :orderName, :attempts, :lastAttempt, :state, :jobSpec, :jobId, :jobResult);
