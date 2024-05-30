package meeting_rooms

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name      string
	intervals [][]int
	expected  bool
}{
	{
		"Test 1",
		[][]int{{0, 30}, {5, 10}, {15, 20}},
		false,
	},
	{
		"Test 2",
		[][]int{{7, 10}, {2, 4}},
		true,
	},
	{
		"Test 3",
		[][]int{{6, 15}, {13, 20}, {6, 17}},
		false,
	},
}

func TestCanAttendMeetings(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := canAttendMeetings_bruteForce(tt.intervals)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
