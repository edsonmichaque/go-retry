package retry

import (
	"time"
)

// Option
type Option func(*Retrier)

// New
func New(opts ...Option) Retrier {
	t := Retrier{}

	t.update(opts...)

	return t
}

// Retrier
type Retrier struct {
	jitter       Jitter
	backoff      Interval
	deadline     time.Duration
	maxRetries   int
	initialDelay time.Duration
}

// WithDeadline
func WithDeadline(d time.Duration) Option {
	return func(t *Retrier) {
		t.deadline = d
	}
}

// WithDelay
func WithDelay(d time.Duration) Option {
	return func(t *Retrier) {
		t.initialDelay = d
	}
}

// WithRetries
func WithRetries(r int) Option {
	return func(t *Retrier) {
		t.maxRetries = r
	}
}

// WithInterval
func WithInterval(b Interval) Option {
	return func(t *Retrier) {
		t.backoff = b
	}
}

// Update
func (t *Retrier) Update(opts ...Option) {
	t.update(opts...)
}

// Clone
func (t *Retrier) Clone() Retrier {
	return *t
}

func (t *Retrier) update(opts ...Option) {
	for _, opt := range opts {
		opt(t)
	}

	if t.jitter == nil {
		WithJitter(0)(t)
	}
}
