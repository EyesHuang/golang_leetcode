package leetcode_730

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type q724 struct {
	param string
	array []int
	ans   int
}

var testCases724 = []q724{
	{
		"scenario 1",
		[]int{1, 7, 3, 6, 5, 6},
		3,
	},
	{
		"scenario 2",
		[]int{1, 2, 3},
		-1,
	},
	{
		"scenario 3",
		[]int{2, 1, -1},
		0,
	},
}

func Test_PivotIndex(t *testing.T) {
	for _, tc := range testCases724 {
		t.Run(tc.param, func(t *testing.T) {
			rsp := PivotIndex(tc.array)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_PivotIndexS2(t *testing.T) {
	for _, tc := range testCases724 {
		t.Run(tc.param, func(t *testing.T) {
			rsp := PivotIndexS2(tc.array)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
