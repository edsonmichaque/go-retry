package retry

import (
	"time"
)

func NewLinear(initial time.Duration) Linear {
	return Linear{
		initial: initial,
	}
}

type Linear struct {
	initial time.Duration
}

func (d Linear) Delay(attempts int) time.Duration {
	if attempts == 0 {
		return 0
	}

	return time.Duration(attempts) * d.initial
}
