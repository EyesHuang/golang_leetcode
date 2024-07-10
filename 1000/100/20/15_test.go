package leetcode_20

import (
	"reflect"
	"testing"
)

var q15TestCases = []struct {
	name     string
	nums     []int
	expected [][]int
}{
	{
		"Test 1",
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	// {
	// 	"Test 2",
	// 	[]int{0, 1, 1},
	// 	[][]int{},
	// },
	// {
	// 	"Test 3",
	// 	[]int{0, 0, 0},
	// 	[][]int{{0, 0, 0}},
	// },
}

func TestThreeSum(t *testing.T) {
	for _, tt := range q15TestCases {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSum(tt.nums)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
