package leetcode_130

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	nums     []int
	expected int
}{
	// {
	// 	"Test 1",
	// 	[]int{100, 4, 200, 1, 3, 2},
	// 	4,
	// },
	// {
	// 	"Test 2",
	// 	[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
	// 	9,
	// },
	{
		"Test 2",
		[]int{1, 2, 0, 1},
		3,
	},
}

func TestLongestConsecutiveHashSet(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestConsecutive_hashSet(tt.nums)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestLongestConsecutiveSorting(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestConsecutive_sorting(tt.nums)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
