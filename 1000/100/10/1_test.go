package leetcode_10

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	nums     []int
	target   int
	expected []int
}{
	{
		"Test 1",
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
	{
		"Test 2",
		[]int{3, 2, 4},
		6,
		[]int{1, 2},
	},
	{
		"Test 3",
		[]int{3, 3},
		6,
		[]int{0, 1},
	},
}

func TestTwoSum_BruteForce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSum_bruteForce(tt.nums, tt.target)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestTwoSum_HashMap(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSum_hashMap(tt.nums, tt.target)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
