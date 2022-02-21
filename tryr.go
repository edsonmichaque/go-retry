package tryr

import (
	"math"
	"time"
)

func Try(max int, fn func() error) {
	var retry bool
	err := fn()
	if err != nil {
		retry = true
	}

	retries := 0
	for retry && retries < max {
		wait := int(math.Exp2(float64(retries))) * 100

		time.Sleep(time.Duration(wait) * time.Millisecond)

		err := fn()
		if err != nil {
			retry = true
		}

		retries += 1
	}
}
