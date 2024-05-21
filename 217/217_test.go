package contains_duplicate

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	nums     []int
	expected bool
}{
	{
		"Test 1",
		[]int{1, 2, 3, 1},
		true,
	},
	{
		"Test 2",
		[]int{1, 2, 3, 4},
		false,
	},
	{
		"Test 1",
		[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		true,
	},
}

func TestContainsDuplicate_BruteForce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsDuplicate_bruteForce(tt.nums)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestContainsDuplicate_set(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsDuplicate_set(tt.nums)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
