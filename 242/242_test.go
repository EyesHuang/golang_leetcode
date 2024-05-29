package valid_anagram

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	s        string
	t        string
	expected bool
}{
	{
		"Test 1",
		"anagram",
		"nagaram",
		true,
	},
	{
		"Test 2",
		"rat",
		"car",
		false,
	},
	{
		"Test 3",
		"ab",
		"a",
		false,
	},
}

func TestIsAnagram_BruteForce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram_bruteForce(tt.s, tt.t)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, waant %v", res, tt.expected)
			}
		})
	}
}
