package retry

import "time"

func New(p Policy) Retryer {
	return Retryer{
		p: p,
	}
}

type Retryer struct {
	maxAttempts int
	p           Policy
}

func (r Retryer) Do(actionFunc func(sequence int) error) Result {
	startedAt := time.Now()

	mustRetry := true

	var totalAttempts int
	for mustRetry && totalAttempts < r.p.MaxAttempts() && time.Since(startedAt) < r.p.Deadline() {
		waitFor := r.p.Delay(totalAttempts)

		if waitFor > 0 {
			time.Sleep(time.Duration(waitFor))
		}

		totalAttempts += 1

		if err := actionFunc(totalAttempts); err == nil {
			mustRetry = false
		}
	}

	return Result{
		TotalAttempts: totalAttempts,
		TotalDuration: time.Since(startedAt),
		Success:       !mustRetry,
	}
}

func WithInitialDelay(p Policy, d time.Duration) Policy {
	return initialDelay{
		Policy: p,
		delay:  d,
	}
}

type initialDelay struct {
	Policy
	delay time.Duration
}

func (w initialDelay) Delay(attempt int) time.Duration {
	if attempt == 0 {
		return w.delay
	}

	return w.Policy.Delay(attempt)
}

func WithMaxAttempts(p Policy, attempts int) Policy {
	return maxAttempts{
		Policy:   p,
		attempts: attempts,
	}
}

type maxAttempts struct {
	Policy
	attempts int
}

func (w maxAttempts) MaxAttempts() int {
	return w.attempts
}
