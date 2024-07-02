package leetcode_260

import (
	"reflect"
	"testing"
)

var q253TestCases = []struct {
	name      string
	intervals [][]int
	expected  int
}{
	{
		"Test 1",
		[][]int{{0, 30}, {5, 10}, {15, 20}},
		2,
	},
	{
		"Test 2",
		[][]int{{7, 10}, {2, 4}},
		1,
	},
}

func TestMinMeetingRooms(t *testing.T) {
	for _, tt := range q253TestCases {
		t.Run(tt.name, func(t *testing.T) {
			res := minMeetingRooms(tt.intervals)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
