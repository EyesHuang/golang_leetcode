package binary_search_test

import (
	"testing"

	q704Bs "leetcode/704"

	"github.com/stretchr/testify/assert"
)

type q704 struct {
	name   string
	array  []int
	target int
	ans    int
}

var testCases = []q704{
	{
		"Find number 9",
		[]int{-1, 0, 3, 5, 9, 12},
		9,
		4,
	},
	{
		"Find number 2",
		[]int{-1, 0, 3, 5, 9, 12},
		2,
		-1,
	},
	{
		"Find number 5",
		[]int{5},
		5,
		0,
	},
}

func Test_Recursion(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rsp := q704Bs.SearchRecursion(tc.array, tc.target)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}

func Test_Loop(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rsp := q704Bs.Search(tc.array, tc.target)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
