package fibonacci_number_test

import (
	"github.com/stretchr/testify/assert"
	q509Fn "leetcode/509"
	"testing"
)

type q509 struct {
	param string
	n     int
	ans   int
}

var testCases = []q509{
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
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q509Fn.FibBruteforce(tc.n)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_FibTopDown(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q509Fn.FibTopDown(tc.n)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_FibDownTop(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q509Fn.FibBottomUp(tc.n)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
