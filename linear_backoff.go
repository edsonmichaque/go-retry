package retry

import (
	"time"
)

func NewLinearBackoff(initial time.Duration) LinearBackoff {
	return LinearBackoff{
		initial: initial,
	}
}

type LinearBackoff struct {
	initial time.Duration
}

func (d LinearBackoff) Delay(attempts int) time.Duration {
	if attempts == 0 {
		return 0
	}

	return time.Duration(attempts) * d.initial
}
