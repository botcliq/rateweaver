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
		// {25, 10 * time.Second},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d*%d", tc.arg1, tc.arg2), func(t *testing.T) {
			r := New(tc.arg1, tc.arg2)
			pv := time.Now()
			t.Logf("the ratelimt duration is : %v", r.(*ratelimit).t)
			lv1 := r.Take()
			t.Logf("The time value returned is: %v,diff %v", lv1, lv1.Sub(pv))
			lv2 := r.Take()
			t.Logf("The time value returned is: %v,diff %v", lv2, lv2.Sub(lv1))
			lv3 := r.Take()
			t.Logf("The time value returned is: %v,diff %v", lv3, lv3.Sub(lv2))
			pv = lv3
			for i := 3; i < 10; i++ {
				cv := r.Take()
				t.Logf("The time value returned is: %v,diff %v", cv, cv.Sub(pv))
				pv = cv
			}

			t.Logf("The time difference are: %v,%v", lv3.Sub(lv2), lv2.Sub(lv1))
			if lv3.Sub(lv2) != lv2.Sub(lv1) {
				t.Errorf("The rate limt is not proper.")
			}
		})
	}
}

// update rateweaver