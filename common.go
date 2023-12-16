package rateweaver

import (
	"time"
)

type ratelimit struct {
	// rate for limiting
	rate int
	t    time.Duration // time to sleep for next execution.
	c    chan time.Time
}

type limiter interface {
	Update(rate int, per time.Duration)
	Take() time.Time
}

func New(rate int, per time.Duration) limiter {
	if rate < 0 || per < 0 {
		return nil
	}
	d := per / time.Duration(rate)
	return &ratelimit{rate: rate, t: d, c: make(chan time.Time)}
}
