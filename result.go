package retry

import "time"

type Result struct {
	Attempts int
	Duration time.Duration
	Success  bool
}
