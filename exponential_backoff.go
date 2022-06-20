package retry

import (
	"time"
)

func NewExponentialBackoff(initial time.Duration) ExponentialBackoff {
	return ExponentialBackoff{
		factor: initial,
	}
}

type ExponentialBackoff struct {
	factor time.Duration
}

func (d ExponentialBackoff) Delay(attempts int) time.Duration {
	if attempts == 0 {
		return 0
	}

	return (1 << (attempts - 1)) * d.factor
}
