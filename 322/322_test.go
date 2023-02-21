package coin_change_test

import (
	"github.com/stretchr/testify/assert"
	q322Fn "leetcode/322"
	"testing"
)

type q322 struct {
	param  string
	coins  []int
	amount int
	ans    int
}

var testCases = []q322{
	{
		"Scenario 1",
		[]int{1, 2},
		3,
		2,
	},
	{
		"Scenario 2",
		[]int{2},
		3,
		-1,
	},
	{
		"Scenario 3",
		[]int{1},
		0,
		0,
	},
}

func Test_CoinChangeBruteForce(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q322Fn.CoinChangeBruteForce(tc.coins, tc.amount)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_CoinChangeTopDown(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q322Fn.CoinChangeTopDown(tc.coins, tc.amount)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
