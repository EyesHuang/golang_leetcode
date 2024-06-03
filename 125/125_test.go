package valid_palindrome

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	input    string
	expected bool
}{
	{
		"Test 1",
		"A man, a plan, a canal: Panama",
		true,
	},
	{
		"Test 2",
		"race a car",
		false,
	},
	{
		"Test 3",
		" ",
		true,
	},
}

func TestIsPalindrome_Reverse(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome_reverse(tt.input)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestIsPalindrome_TwoPointers(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome_twoPointers(tt.input)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
