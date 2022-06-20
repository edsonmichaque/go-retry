package retry

import (
	"time"
)

func NewFixed(initial time.Duration) Fixed {
	return Fixed{
		interval: initial,
	}
}

type Fixed struct {
	interval time.Duration
}

func (d Fixed) Delay(attempts int) time.Duration {
	if attempts == 0 {
		return 0
	}

	return d.interval
}
