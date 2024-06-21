package leetcode_240

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	nums     []int
	expected []int
}{
	{
		"Test 1",
		[]int{1, 2, 3, 4},
		[]int{24, 12, 8, 6},
	},
	{
		"Test 2",
		[]int{-1, 1, 0, -3, 3},
		[]int{0, 0, 9, 0, 0},
	},
}

func TestProductExceptSelf(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := productExceptSelf(tt.nums)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestProductExceptSelf_Optimization(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := productExceptSelf_optimization(tt.nums)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
