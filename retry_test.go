package retry_test

import (
	"testing"
	"time"

	"gitlab.com/edsonmichaque/go-retry"
)

func Test(t *testing.T) {

	r := retry.New(retry.WithMaxAttempts(retry.WithInitialDelay(retry.NewExponential(100*time.Millisecond), 5*time.Millisecond), 10))

	res := r.Do(func(int) error {
		return nil
	})

	_ = res
}
