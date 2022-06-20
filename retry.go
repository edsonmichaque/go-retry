package retry

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"
)

func New(p Backoff) Retry {
	return Retry{
		delayer: p,
	}
}

type Backoff interface {
	Delay(int) time.Duration
}

type AttemptsLimiter interface {
	MaxAttempts() int
}

type DeadlineLimiter interface {
	Deadline() time.Duration
}

type Retry struct {
	delayer Backoff
}

func (r Retry) Do(callback func(sequence int) error) Result {
	startingTime := time.Now()

	retry := true

	maxAttempts := math.MaxInt
	if limiter, ok := r.delayer.(AttemptsLimiter); ok {
		maxAttempts = limiter.MaxAttempts()
	}

	var deadline time.Duration = 1<<61 - 1
	if limiter, ok := r.delayer.(DeadlineLimiter); ok {
		deadline = limiter.Deadline()
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

func WithInitialDelay(p Backoff, d time.Duration) Backoff {
	return initialDelay{
		Backoff: p,
		delay:   d,
	}
}

type initialDelay struct {
	Backoff
	delay time.Duration
}

func (j initialDelay) Deadline() time.Duration {
	return DefaultDeadline(j.Backoff)
}

func (j initialDelay) MaxAttempts() time.Duration {
	return DefaultDeadline(j.Backoff)
}

func (w initialDelay) Delay(attempt int) time.Duration {
	if attempt == 0 {
		return w.delay
	}

	return w.Backoff.Delay(attempt)
}

func WithMaxAttempts(p Backoff, attempts int) Backoff {
	return maxAttempts{
		Backoff:  p,
		attempts: attempts,
	}
}

type maxAttempts struct {
	Backoff
	attempts int
}

func (w maxAttempts) MaxAttempts() int {
	return w.attempts
}

func (j maxAttempts) Deadline() time.Duration {
	return DefaultDeadline(j.Backoff)
}

type deadline struct {
	Backoff
	value time.Duration
}

func (d deadline) Deadline() time.Duration {
	return d.value
}

func (j deadline) MaxAttempts() time.Duration {
	return DefaultMaxAttempts(j.Backoff)
}

func WithDeadline(b Backoff, v time.Duration) Backoff {
	return deadline{
		Backoff: b,
		value:   v,
	}
}

type jitter struct {
	Backoff
}

func (j jitter) Delay(attempts int) time.Duration {
	if attempts == 0 {
		return 0
	}

	delay, err := rand.Int(rand.Reader, big.NewInt(int64(j.Backoff.Delay(attempts))))
	if err != nil {
		return 0
	}

	return time.Duration(delay.Int64())
}

func (j jitter) Deadline() time.Duration {
	return DefaultDeadline(j.Backoff)
}

func (j jitter) MaxAttempts() time.Duration {
	return DefaultDeadline(j.Backoff)
}

func WithJitter(b Backoff) Backoff {
	return jitter{
		Backoff: b,
	}
}

func DefaultMaxAttempts(b Backoff) time.Duration {
	if v, ok := b.(AttemptsLimiter); ok {
		return time.Duration(v.MaxAttempts())
	}

	return math.MaxInt

}

func DefaultDeadline(b Backoff) time.Duration {
	if v, ok := b.(DeadlineLimiter); ok {
		return time.Duration(v.Deadline())
	}

	return (1 << 63) - 1

}
