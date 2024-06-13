package leetcode_250

import (
	"reflect"
	"testing"
)

var tests242 = []struct {
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
	{
		"Test 4",
		"acca",
		"acaa",
		false,
	},
}

func TestIsAnagram_BruteForce(t *testing.T) {
	for _, tt := range tests242 {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram_bruteForce(tt.s, tt.t)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, waant %v", res, tt.expected)
			}
		})
	}
}

func TestIsAnagram_FrequencyCounter(t *testing.T) {
	for _, tt := range tests242 {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram_frequencyCounter(tt.s, tt.t)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestIsAnagram_Sorting(t *testing.T) {
	for _, tt := range tests242 {
		t.Run(tt.name, func(t *testing.T) {
			res := isAnagram_sorting(tt.s, tt.t)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
