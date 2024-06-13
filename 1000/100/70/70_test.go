package leetcode_70

import "testing"

var tests = []struct {
	name     string
	stair    int
	expected int
}{
	{
		"Test 1",
		2,
		2,
	},
	{
		"Test 2",
		3,
		3,
	},
	{
		"Test 3",
		5,
		8,
	},
}

func TestClimbStairs_BruteForce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ways := climbStairs_BruteForce(tt.stair)

			if ways != tt.expected {
				t.Errorf("got %d, want %d", ways, tt.expected)
			}
		})
	}
}

func TestClimbStairs_Memo(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ways := climbStairs_Memo(tt.stair)

			if ways != tt.expected {
				t.Errorf("got %d, want %d", ways, tt.expected)
			}
		})
	}
}

func TestClimbStairs_DP(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ways := climbStairs_DP(tt.stair)

			if ways != tt.expected {
				t.Errorf("got %d, want %d", ways, tt.expected)
			}
		})
	}
}
