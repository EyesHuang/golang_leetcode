package leetcode_20

import (
	"reflect"
	"testing"
)

var q11TestCases = []struct {
	name     string
	height   []int
	expected int
}{
	{
		"Test 1",
		[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		49,
	},
	{
		"Test 2",
		[]int{1, 1},
		1,
	},
}

func TestMaxAreaBruteForce(t *testing.T) {
	for _, tt := range q11TestCases {
		t.Run(tt.name, func(t *testing.T) {
			res := maxArea_bruteForce(tt.height)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestMaxAreaTwoPointers(t *testing.T) {
	for _, tt := range q11TestCases {
		t.Run(tt.name, func(t *testing.T) {
			res := maxArea_twoPointers(tt.height)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
