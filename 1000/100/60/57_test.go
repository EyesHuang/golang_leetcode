package leetcode_60

import (
	"reflect"
	"testing"
)

var q56TestCases = []struct {
	name        string
	intervals   [][]int
	newInterval []int
	expected    [][]int
}{
	{
		"Test 1",
		[][]int{{1, 3}, {6, 9}},
		[]int{2, 5},
		[][]int{{1, 5}, {6, 9}},
	},
	{
		"Test 2",
		[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
		[]int{4, 8},
		[][]int{{1, 2}, {3, 10}, {12, 16}},
	},
	{
		"Test 3",
		[][]int{{1, 5}},
		[]int{6, 8},
		[][]int{{1, 5}, {6, 8}},
	},
	{
		"Test 4",
		[][]int{{1, 5}},
		[]int{0, 3},
		[][]int{{0, 5}},
	},
}

func TestInsertLinear(t *testing.T) {
	for _, tt := range q56TestCases {
		t.Run(tt.name, func(t *testing.T) {
			res := insert_linear(tt.intervals, tt.newInterval)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestInsertBinarySearch(t *testing.T) {
	for _, tt := range q56TestCases {
		t.Run(tt.name, func(t *testing.T) {
			res := insert_binarySearch(tt.intervals, tt.newInterval)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
