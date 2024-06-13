package leetcode_1480

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type q1480 struct {
	param string
	array []int
	ans   []int
}

var testCases1480 = []q1480{
	{
		"scenario 1",
		[]int{1, 2, 3, 4},
		[]int{1, 3, 6, 10},
	},
	{
		"scenario 2",
		[]int{1, 1, 1, 1, 1},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"scenario 3",
		[]int{3, 1, 2, 10, 1},
		[]int{3, 4, 6, 16, 17},
	},
}

func Test_RunningSum(t *testing.T) {
	for _, tc := range testCases1480 {
		t.Run(tc.param, func(t *testing.T) {
			rsp := runningSum(tc.array)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
