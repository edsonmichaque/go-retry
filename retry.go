package retry

import "time"

// Retry
func (t Retrier) Retry(fn func(sequence int) error, opts ...Option) Report {
	t.update(opts...)

	startedAt := time.Now()

	if t.initialDelay > 0 {
		time.Sleep(t.initialDelay)
	}

	attempts := 1
	shouldRetry := true

	if err := fn(0); err != nil {
		for shouldRetry && attempts < t.maxRetries &&
			(time.Since(startedAt) < t.deadline || t.deadline == time.Duration(0)) {
			waitingInterval := t.backoff(attempts-1)*100 + t.jitter()

			time.Sleep(time.Duration(waitingInterval) * time.Millisecond)

			if err := fn(attempts); err == nil {
				shouldRetry = false
			}

			attempts += 1
		}
	}

	return Report{
		TotalAttempts: attempts,
		TotalDuration: time.Since(startedAt),
		Success:       !shouldRetry,
	}
}
