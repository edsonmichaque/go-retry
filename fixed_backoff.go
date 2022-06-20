package retry

import (
	"time"
)

func NewFixedBackoff(initial time.Duration) FixedBackoff {
	return FixedBackoff{
		interval: initial,
	}
}

type FixedBackoff struct {
	interval time.Duration
}

func (d FixedBackoff) Delay(attempts int) time.Duration {
	if attempts == 0 {
		return 0
	}

	return d.interval
}
