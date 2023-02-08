package running_sum_of_1d_array_test

import (
	"github.com/stretchr/testify/assert"
	q1480Rs1at "leetcode/1480"
	"testing"
)

type q1480 struct {
	param string
	array []int
	ans   []int
}

var testCases = []q1480{
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
	for _, tc := range testCases {
		t.Run(tc.param, func(t *testing.T) {
			rsp := q1480Rs1at.RunningSum(tc.array)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
