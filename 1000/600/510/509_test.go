package leetcode_510

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type q509 struct {
	param string
	n     int
	ans   int
}

var testCases509 = []q509{
	{
		"scenario 1",
		2,
		1,
	},
	{
		"scenario 2",
		3,
		2,
	},
	{
		"scenario 3",
		4,
		3,
	},
}

func Test_FibBruteForce(t *testing.T) {
	for _, tc := range testCases509 {
		t.Run(tc.param, func(t *testing.T) {
			rsp := FibBruteforce(tc.n)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_FibTopDown(t *testing.T) {
	for _, tc := range testCases509 {
		t.Run(tc.param, func(t *testing.T) {
			rsp := FibTopDown(tc.n)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_FibBottomUp(t *testing.T) {
	for _, tc := range testCases509 {
		t.Run(tc.param, func(t *testing.T) {
			rsp := FibBottomUp(tc.n)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_FibTwoVars(t *testing.T) {
	for _, tc := range testCases509 {
		t.Run(tc.param, func(t *testing.T) {
			rsp := FibTwoVars(tc.n)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
