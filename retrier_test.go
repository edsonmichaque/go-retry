package retry_test

import (
	"errors"
	"testing"
	"time"

	"gitlab.com/edsonmichaque/go-retry"
)

func TestRetrier(t *testing.T) {
	retry := retry.New(
		retry.WithInterval(retry.Binary()),
		retry.WithDeadline(5*time.Second),
		retry.WithJitter(1000),
		retry.WithRetries(16),
	)

	stats := retry.Retry(func(attempts int) error {
		t.Log("attempt: ", attempts)

		return errors.New("s")
	})

	t.Logf("Stats: %#v", stats)
}
