package retry

import "time"

// Report
type Report struct {
	TotalAttempts int
	TotalDuration time.Duration
	Success       bool
}
