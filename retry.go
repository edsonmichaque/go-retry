package retry

import (
	"math"
	"time"
)

func New(p Delayer) Retryer {
	return Retryer{
		delayer: p,
	}
}

type Delayer interface {
	Delay(int) time.Duration
}

type AttemptsLimiter interface {
	MaxAttempts() int
}

type TimeLimiter interface {
	Deadline() time.Duration
}

type Retryer struct {
	delayer Delayer
}

func (r Retryer) Do(callback func(sequence int) error) Result {
	startingTime := time.Now()

	retry := true

	maxAttempts := math.MaxInt
	if attemptsLimiter, ok := r.delayer.(AttemptsLimiter); ok {
		maxAttempts = attemptsLimiter.MaxAttempts()
	}

	deadline := time.Minute
	if timeLimiter, ok := r.delayer.(TimeLimiter); ok {
		deadline = timeLimiter.Deadline()
	}

	var attempts int
	for retry && attempts < maxAttempts && time.Since(startingTime) < deadline {
		delay := r.delayer.Delay(attempts)

		if delay > 0 {
			time.Sleep(time.Duration(delay))
		}

		if err := callback(attempts); err == nil {
			retry = false
		}

		attempts += 1

	}

	duration := time.Since(startingTime)

	return Result{
		Attempts: attempts,
		Duration: duration,
		Success:  !retry,
	}
}

func WithInitialDelay(p Delayer, d time.Duration) Delayer {
	return initialDelay{
		Delayer: p,
		delay:   d,
	}
}

type initialDelay struct {
	Delayer
	delay time.Duration
}

func (w initialDelay) Delay(attempt int) time.Duration {
	if attempt == 0 {
		return w.delay
	}

	return w.Delayer.Delay(attempt)
}

func WithMaxAttempts(p Delayer, attempts int) Delayer {
	return maxAttempts{
		Delayer:  p,
		attempts: attempts,
	}
}

type maxAttempts struct {
	Delayer
	attempts int
}

func (w maxAttempts) MaxAttempts() int {
	return w.attempts
}
