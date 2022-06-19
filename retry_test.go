package retry_test

import (
	"testing"
	"time"

	"gitlab.com/edsonmichaque/go-retry"
	"gitlab.com/edsonmichaque/go-retry/exponential"
)

func Test(t *testing.T) {
	var p retry.Policy
	p = exponential.NewPolicy(
		exponential.WithDeadline(5 * time.Second),
	)

	p = retry.WithInitialDelay(p, 5*time.Second)

	r := retry.New(p)

	res := r.Do(func(int) error {
		return nil
	})

	_ = res
}
