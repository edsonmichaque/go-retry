package retry

import (
	"time"
)

type Policy interface {
	Delay(int) time.Duration
	MaxAttempts() int
	Deadline() time.Duration
}
