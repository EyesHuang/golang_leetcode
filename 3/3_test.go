package longest_substring_without_repeating_characters_test

import (
	"testing"

	q3Lswrc "leetcode/3"

	"github.com/stretchr/testify/assert"
)

type q3 struct {
	param string
	ans   int
}

var testCases = []q3{
	{
		"abcabcbb",
		3,
	},
	{
		"bbbb",
		1,
	},
	{
		"pwwkew",
		3,
	},
	{
		"dvdf",
		3,
	},
	{
		"abba",
		2,
	},
}

func Test_BruteForce(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q3Lswrc.LengthOfLongestSubstringBruteForce(tc.param)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_S1(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q3Lswrc.LengthOfLongestSubstringS1(tc.param)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_S2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q3Lswrc.LengthOfLongestSubstring(tc.param)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
