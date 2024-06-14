package leetcode_50

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			input: []string{""},
			expected: [][]string{
				{""},
			},
		},
		{
			input: []string{"a"},
			expected: [][]string{
				{"a"},
			},
		},
		{
			input: []string{"", "", ""},
			expected: [][]string{
				{"", "", ""},
			},
		},
		{
			input: []string{"abc", "bca", "cab", "bac", "acb", "cba"},
			expected: [][]string{
				{"abc", "bca", "cab", "bac", "acb", "cba"},
			},
		},
		{
			input: []string{"listen", "silent", "enlist", "google", "gogole", "gooegl"},
			expected: [][]string{
				{"listen", "silent", "enlist"},
				{"google", "gogole", "gooegl"},
			},
		},
	}

	for _, test := range tests {
		result := groupAnagrams(test.input)
		if !compareGroups(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}

func compareGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count occurrences of each group
	mapA := make(map[string]int)
	mapB := make(map[string]int)

	for _, group := range a {
		key := createGroupKey(group)
		mapA[key]++
	}

	for _, group := range b {
		key := createGroupKey(group)
		mapB[key]++
	}

	return reflect.DeepEqual(mapA, mapB)
}

func createGroupKey(group []string) string {
	// Create a unique key for each group by sorting the strings within the group
	sortedGroup := make([]string, len(group))
	copy(sortedGroup, group)
	sort.Strings(sortedGroup)
	return strings.Join(sortedGroup, ",")
}
