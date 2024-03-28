package merge_two_sorted_lists

import (
	"leetcode/util"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var emptySlice []int

	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "Test 1: Merge two non-empty lists",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Test 2: Merge one empty list and one non-empty list",
			list1:    []int{},
			list2:    []int{1, 3, 4},
			expected: []int{1, 3, 4},
		},
		{
			name:     "Test 3: Merge two empty lists",
			list1:    []int{},
			list2:    []int{},
			expected: emptySlice,
		},
		{
			name:     "Test 4: Merge lists with equal elements",
			list1:    []int{2, 2, 2},
			list2:    []int{2, 2, 2},
			expected: []int{2, 2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := util.CreateLinkedList(tt.list1)
			l2 := util.CreateLinkedList(tt.list2)

			result := mergeTwoLists(l1, l2)
			resultSlice := util.LinkedListToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("mergeTwoLists(%v, %v) got %v, want %v", tt.list1, tt.list2, resultSlice, tt.expected)
			}
		})
	}
}
