package leetcode_210

import (
	"reflect"
	"testing"

	"leetcode/util"
)

var (
	emptySlice206 []int

	tests206 = []struct {
		name     string
		head     []int
		expected []int
	}{
		{
			"Test 1",
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			"Test 2",
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			"Test 3",
			[]int{},
			emptySlice206,
		},
	}
)

func TestReverseListIterative(t *testing.T) {
	for _, tt := range tests206 {
		t.Run(tt.name, func(t *testing.T) {
			headList := util.CreateLinkedList(tt.head)

			res := reverseList_iterative(headList)
			resSlice := util.LinkedListToSlice(res)

			if !reflect.DeepEqual(resSlice, tt.expected) {
				t.Errorf("got %v, want %v", resSlice, tt.expected)
			}
		})
	}
}

func TestReverseListRecursive(t *testing.T) {
	for _, tt := range tests206 {
		t.Run(tt.name, func(t *testing.T) {
			headList := util.CreateLinkedList(tt.head)

			res := reverseList_recursive(headList)
			resSlice := util.LinkedListToSlice(res)

			if !reflect.DeepEqual(resSlice, tt.expected) {
				t.Errorf("got %v, want %v", resSlice, tt.expected)
			}
		})
	}
}
