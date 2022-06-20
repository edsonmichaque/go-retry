package retry

import (
	"math"
	"time"
)

func NewExponentialBackoff(initial time.Duration) ExponentialBackoff {
	return ExponentialBackoff{
		Initial: initial,
	}
}

type ExponentialBackoff struct {
	Initial time.Duration
}

func (d ExponentialBackoff) Delay(attempts int) time.Duration {
	if attempts == 0 {
		return 0
	}

	return time.Duration(int(math.Pow(2, float64(attempts-1)))) * d.Initial
}
