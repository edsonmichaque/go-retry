package retry_test

import (
	"errors"
	"testing"
	"time"

	"gitlab.com/edsonmichaque/go-retry"
)

func TestRetry(t *testing.T) {

	tcases := map[string]struct {
		interval retry.Interval
		deadline time.Duration
		retries  int
		jitter   int
	}{
		"binary backoff": {
			interval: retry.Binary(),
		},
		"linear backoff": {
			interval: retry.Binary(),
		},
	}

	for tname, tcase := range tcases {
		t.Run(tname, func(t *testing.T) {
			r := retry.New(
				retry.WithInterval(tcase.interval),
				retry.WithDeadline(tcase.deadline),
				retry.WithJitter(tcase.jitter),
				retry.WithRetries(tcase.retries),
			)

			_ = r.Retry(func(attempts int) error {
				t.Log("attempt: ", attempts)

				return errors.New("s")
			})
		})
	}
}
