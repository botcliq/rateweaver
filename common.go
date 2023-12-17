/*
 *   Copyright (c) 2023
 *   All rights reserved.
 */
package rateweaver

import (
	"fmt"
	"time"
)

type ratelimit struct {
	// rate for limiting
	rate int
	t    time.Duration // time to sleep for next execution.
	c    chan time.Time
}

type Limiter interface {
	Update(rate int, per time.Duration)
	Take() time.Time
}

func New(rate int, per time.Duration) Limiter {
	if rate < 0 || per < 0 {
		return nil
	}
	d := per / time.Duration(rate)
	fmt.Sprintf("The Duration for sleep is : %v", d)
	r := ratelimit{rate: rate, t: d, c: make(chan time.Time)}
	go r.start()
	return &r
}
