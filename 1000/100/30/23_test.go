package leetcode_30

import (
	"reflect"
	"testing"

	"leetcode/util"
)

var (
	emptySlice23 []int

	tests23 = []struct {
		name     string
		list     [][]int
		expected []int
	}{
		{
			"Test 1",
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			"Test 2",
			[][]int{},
			emptySlice23,
		},
		{
			"Test 3",
			[][]int{{}},
			emptySlice23,
		},
	}
)

func TestMergeKLists_BruteForce(t *testing.T) {
	for _, tt := range tests23 {
		var lists []*util.ListNode

		for _, l := range tt.list {
			l1 := util.CreateLinkedList(l)
			lists = append(lists, l1)
		}

		result := mergeKLists_BruteForce(lists)
		resultSlice := util.LinkedListToSlice(result)

		if !reflect.DeepEqual(tt.expected, resultSlice) {
			t.Errorf("got %v, want %v", resultSlice, tt.expected)
		}
	}
}
