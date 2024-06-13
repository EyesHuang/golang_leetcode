package leetcode_10

import (
	"testing"

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
			rsp := LengthOfLongestSubstringBruteForce(tc.param)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_S1(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := LengthOfLongestSubstringS1(tc.param)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_S2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := LengthOfLongestSubstringS2(tc.param)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
