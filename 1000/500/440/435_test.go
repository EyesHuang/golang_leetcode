package leetcode_440

import (
	"reflect"
	"testing"
)

var q435TestCases = []struct {
	name      string
	intervals [][]int
	expected  int
}{
	{
		"Test 1",
		[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		1,
	},
	{
		"Test 2",
		[][]int{{1, 2}, {1, 2}, {1, 2}},
		2,
	},
	{
		"Test 3",
		[][]int{{1, 2}, {2, 3}},
		0,
	},
}

func TestEraseOverlapIntervals(t *testing.T) {
	for _, tt := range q435TestCases {
		t.Run(tt.name, func(t *testing.T) {
			res := eraseOverlapIntervals(tt.intervals)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
