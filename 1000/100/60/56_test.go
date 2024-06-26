package leetcode_60

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name      string
	intervals [][]int
	expected  [][]int
}{
	{
		"Test 1",
		[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		[][]int{{1, 6}, {8, 10}, {15, 18}},
	},
	{
		"Test 2",
		[][]int{{1, 4}, {4, 5}},
		[][]int{{1, 5}},
	},
	{
		"Test 31",
		[][]int{{1, 4}, {2, 3}},
		[][]int{{1, 4}},
	},
}

func TestMerge(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := merge(tt.intervals)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
