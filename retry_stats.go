package retry

import "time"

// Stats
type Stats struct {
	retries  int
	duration time.Duration
	success  bool
}
