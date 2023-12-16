package rateweaver

import "time"

func (r *ratelimit) Update(rate int, per time.Duration) {

}

func (r *ratelimit) Take() time.Time {
	return time.Now()
}
