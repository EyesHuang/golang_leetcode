package leetcode_130

import (
	"reflect"
	"testing"
)

var tests121 = []struct {
	name     string
	prices   []int
	expected int
}{
	{
		"Test 1",
		[]int{7, 1, 5, 3, 6, 4},
		5,
	},
	{
		"Test 2",
		[]int{7, 6, 4, 3, 1},
		0,
	},
}

func TestMaxProfit_BruteForce(t *testing.T) {
	for _, tt := range tests121 {
		t.Run(tt.name, func(t *testing.T) {
			res := maxProfit_bruteForce(tt.prices)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestMaxProfit_SlidingWindow(t *testing.T) {
	for _, tt := range tests121 {
		t.Run(tt.name, func(t *testing.T) {
			res := maxProfit_slidingWindow(tt.prices)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
