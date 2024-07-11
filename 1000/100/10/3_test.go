package leetcode_10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var q3TestCases = []struct {
	param string
	ans   int
}{
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

func TestLengthOfLongestSubstring_BruteForce(t *testing.T) {
	for _, tc := range q3TestCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := lengthOfLongestSubstring_bruteForce(tc.param)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func TestLengthOfLongestSubstring_SlidingWindow(t *testing.T) {
	for _, tc := range q3TestCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := lengthOfLongestSubstring_slidingWindow(tc.param)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
