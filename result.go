package retry

import "time"

type Result struct {
	TotalAttempts int
	TotalDuration time.Duration
	Success       bool
}
