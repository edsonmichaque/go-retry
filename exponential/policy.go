package exponential

import "time"

type Option func(*policy)

func NewPolicy(opts ...Option) policy {
	p := policy{}

	for _, opt := range opts {
		opt(&p)
	}

	return p
}

func WithDeadline(d time.Duration) Option {
	return func(p *policy) {
		p.deadline = d
	}
}

type policy struct {
	interval     time.Duration
	initialDelay time.Duration
	maxAttempts  int
	deadline     time.Duration
}

func (p policy) Delay(current int) time.Duration {
	return time.Duration(int(p.interval) * current * 2)
}

func (p policy) MaxAttempts() int {
	return 0
}

func (p policy) Deadline() time.Duration {
	return 60 * time.Second
}
