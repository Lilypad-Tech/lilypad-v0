INSERT INTO events
	(orderId, orderOwner, orderNumber, orderResultType, attempts, lastAttempt, state, jobSpec, jobId, jobResult, jobStdout, jobStderr, jobExitcode)
    VALUES (:orderId, :orderOwner, :orderNumber, :orderResultType, :attempts, :lastAttempt, :state, :jobSpec, :jobId, :jobResult, :jobStdout, :jobStderr, :jobExitcode);
