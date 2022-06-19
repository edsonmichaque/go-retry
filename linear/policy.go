package exponential

import "time"

type Option func(*Policy)

func WithInterval(d time.Duration) Option {
	return func(p *Policy) {
		p.interval = d
	}
}

func WithInitialDelay(d time.Duration) Option {
	return func(p *Policy) {
		p.initialDelay = d
	}
}

func WithMaxAttempts(max int) Option {
	return func(p *Policy) {
		p.maxAttempts = max
	}
}

func WithDeadline(d time.Duration) Option {
	return func(p *Policy) {
		p.deadline = d
	}
}

func NewPolicy(opts ...Option) Policy {
	p := Policy{}

	for _, opt := range opts {
		opt(&p)
	}

	return p
}

type Policy struct {
	interval     time.Duration
	initialDelay time.Duration
	maxAttempts  int
	deadline     time.Duration
}

func (e Policy) Backoff(current int) time.Duration {
	return time.Duration(int(e.interval) * current * 2)
}

func (e Policy) InitialDelay() time.Duration {
	return 0
}

func (e Policy) MaxAttempts() int {
	return 0
}

func (e Policy) Deadline() time.Duration {
	return 0
}

func (e Policy) Jitter(max int) time.Duration {
	return 0
}
