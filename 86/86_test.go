package partition_list

import (
	"reflect"
	"testing"

	"leetcode/util"
)

func TestPartitionList(t *testing.T) {
	tests := []struct {
		name     string
		head     []int
		x        int
		expected []int
	}{
		{
			"Test Case 1",
			[]int{1, 4, 3, 2, 5, 2},
			3,
			[]int{1, 2, 2, 4, 3, 5},
		},
		{
			"Test Case 2",
			[]int{2, 1},
			2,
			[]int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := util.CreateLinkedList(tt.head)
			result := partition(list, tt.x)
			resultSlice := util.LinkedListToSlice(result)

			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("got %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}
