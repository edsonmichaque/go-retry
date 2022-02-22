package retry

import (
	"math/rand"
	"time"
)

type Jitter func() int

func WithJitter(max int) Option {
	return func(t *Retrier) {
		t.jitter = func() int {
			if max == 0 {
				return max
			}

			rand.Seed(time.Now().UnixNano())
			return rand.Intn(max)
		}
	}
}
