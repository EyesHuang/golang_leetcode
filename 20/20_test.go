package valid_parentheses

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
		"()",
		true,
	},
	{
		"Test 2",
		"()[]{}",
		true,
	},
	{
		"Test 3",
		"(]",
		false,
	},
	{
		"Test 4",
		"){",
		false,
	},
	{
		"Test 5",
		"(((((((()",
		false,
	},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isValid(tt.input)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
