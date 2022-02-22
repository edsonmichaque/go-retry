package retry

import "time"

// Retry
func (t Retrier) Retry(fn func(sequence int) error, opts ...Option) Stats {
	t.update(opts...)

	startedAt := time.Now()

	if t.delay > 0 {
		time.Sleep(t.delay)
	}

	attempts := 1
	shouldRetry := true

	if err := fn(0); err != nil {
		for shouldRetry && attempts < t.retries &&
			(time.Since(startedAt) < t.timeout || t.timeout == time.Duration(0)) {
			waitingInterval := t.backoff(attempts-1)*100 + t.jitter()

			time.Sleep(time.Duration(waitingInterval) * time.Millisecond)

			if err := fn(attempts); err == nil {
				shouldRetry = false
			}

			attempts += 1
		}
	}

	return Stats{
		retries:  attempts,
		duration: time.Since(startedAt),
		success:  !shouldRetry,
	}
}
