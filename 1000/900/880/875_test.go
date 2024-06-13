package leetcode_880

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type q875 struct {
	name   string
	array  []int
	target int
	ans    int
}

var testCases875 = []q875{
	{
		"Scenario 1",
		[]int{3, 6, 7, 11},
		8,
		4,
	},
	{
		"Scenario 2",
		[]int{30, 11, 23, 4, 20},
		5,
		30,
	},
	{
		"Scenario 3",
		[]int{30, 11, 23, 4, 20},
		6,
		23,
	},
}

func Test_MinEatingSpeed(t *testing.T) {
	for _, tc := range testCases875 {
		t.Run(tc.name, func(t *testing.T) {
			rsp := minEatingSpeed(tc.array, tc.target)
			assert.Equal(t, tc.ans, rsp)
		})
	}
}
