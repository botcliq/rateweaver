/*
 *   Copyright (c) 2023
 *   All rights reserved.
 */
package rateweaver

import (
	"fmt"
	"time"
)

func (r *ratelimit) Update(rate int, per time.Duration) {
	d := per / time.Duration(rate)
	fmt.Sprintf("The New Duration for sleep is : %v", d)
	r.t = d
}

func (r *ratelimit) Take() time.Time {
	return <-r.c
}

func (r *ratelimit) start() {
	for {
		time.Sleep(r.t)
		r.c <- time.Now()
	}
}
