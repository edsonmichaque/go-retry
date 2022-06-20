package retry_test

import (
	"errors"
	"testing"
	"time"

	"gitlab.com/edsonmichaque/go-retry"
)

func Test32(t *testing.T) {

	var d retry.Backoff

	d = retry.NewExponentialBackoff(1000 * time.Millisecond)
	d = retry.WithInitialDelay(d, 5*time.Second)
	d = retry.WithDeadline(d, 60*time.Second)
	d = retry.WithJitter(d)
	d = retry.WithMaxAttempts(d, 20)

	r := retry.New(d)
	var attempt int
	res := r.Do(func(a int) error {
		t.Logf("attempt %d", attempt+1)

		attempt += 1
		time.Sleep(100 * time.Millisecond)

		return errors.New("error")
	})

	_ = res
}
