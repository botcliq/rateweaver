/*
 *   Copyright (c) 2023
 *   All rights reserved.
 */
package rateweaver

import (
	"fmt"
	"testing"
	"time"
)

type testCase struct {
	arg1 int
	arg2 time.Duration
}

type Update struct {
	arg1 int
	arg2 time.Duration
	arg3 time.Duration
}

// Create rateweaver
func TestCreate(t *testing.T) {
	cases := []testCase{
		{1, 2 * time.Second},
		{25, 10 * time.Second},
		{1000, 1 * time.Second},
		{2, 2 * time.Minute},
		{200, 2 * time.Minute},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d*%d", tc.arg1, tc.arg2), func(t *testing.T) {
			r := New(tc.arg1, tc.arg2)
			if r == nil {
				t.Errorf("Unable to create Limiter: Expected object got :%v", r)
			}
		})
	}
}

// take rateweaver
func TestTake(t *testing.T) {
	cases := []testCase{
		{10, 1 * time.Second},
		{10, 5 * time.Second},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d*%d", tc.arg1, tc.arg2), func(t *testing.T) {
			r := New(tc.arg1, tc.arg2)
			pv := time.Now()
			var td time.Duration
			for i := 3; i < 10; i++ {
				cv := r.Take()
				td += cv.Sub(pv)
				t.Logf("The time value returned is: %v,diff %v", cv, cv.Sub(pv))
				pv = cv
			}
			if td/r.(*ratelimit).t == 1 {
				t.Errorf("The rate limt is not proper.")
			}
		})
	}
}

// update rateweaver
func TestUpdate(t *testing.T) {
	cases := []Update{
		{10, 1 * time.Second, 2 * time.Second},
		{10, 5 * time.Second, 1 * time.Second},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d*%d", tc.arg1, tc.arg2), func(t *testing.T) {
			r := New(tc.arg1, tc.arg2)
			pv := time.Now()
			var td time.Duration
			for i := 3; i < 10; i++ {
				cv := r.Take()
				td += cv.Sub(pv)
				t.Logf("The time value returned is: %v,diff %v", cv, cv.Sub(pv))
				pv = cv
			}
			if td/r.(*ratelimit).t == 1 {
				t.Errorf("The rate limt is not proper.")
			}
			r.Update(tc.arg1, tc.arg3)
			pv = time.Now()
			td = 0
			for i := 3; i < 10; i++ {
				cv := r.Take()
				td += cv.Sub(pv)
				t.Logf("The updated time value returned is: %v,diff %v", cv, cv.Sub(pv))
				pv = cv
			}
			if td/r.(*ratelimit).t == 1 {
				t.Errorf("The updated rate limt is not proper.")
			}
		})
	}
}
