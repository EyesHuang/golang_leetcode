package leetcode_350

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	nums     []int
	k        int
	expected []int
}{
	{
		"Test 1",
		[]int{1, 1, 1, 2, 2, 3},
		2,
		[]int{1, 2},
	},
	{
		"Test 2",
		[]int{1},
		1,
		[]int{1},
	},
	{
		"Test 3",
		[]int{5, 3, 1, 1, 1, 3, 73, 1},
		2,
		[]int{1, 3},
	},
	{
		"Test 4",
		[]int{1, 2},
		2,
		[]int{1, 2},
	},
}

func TestTopKFrequent(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := topKFrequent(tt.nums, tt.k)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestTopKFrequentHeap(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := topKFrequent_heap(tt.nums, tt.k)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
