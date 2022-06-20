package retry

import (
	"math"
	"time"
)

func NewExponential(initial time.Duration) Exponential {
	return Exponential{
		Initial: initial,
	}
}

type Exponential struct {
	Initial time.Duration
}

func (d Exponential) Delay(attempts int) time.Duration {
	if attempts == 0 {
		return 0
	}

	return time.Duration(int(math.Pow(2, float64(attempts-1)))) * d.Initial
}
