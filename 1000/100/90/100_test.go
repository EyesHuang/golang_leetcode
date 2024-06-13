package leetcode_90

import (
	"reflect"
	"testing"

	"leetcode/util"
)

var tests = []struct {
	name     string
	p        *util.TreeNode
	q        *util.TreeNode
	expected bool
}{
	{
		name: "Test 1",
		p: &util.TreeNode{
			Val: 1,
			Left: &util.TreeNode{
				Val:   2,
				Left:  &util.TreeNode{Val: 1},
				Right: &util.TreeNode{Val: 3},
			},
			Right: nil,
		},
		q: &util.TreeNode{
			Val:  1,
			Left: nil,
			Right: &util.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		expected: false,
	},
	{
		name: "Test 2",
		p: &util.TreeNode{
			Val:   1,
			Left:  &util.TreeNode{Val: 2},
			Right: &util.TreeNode{Val: 3},
		},
		q: &util.TreeNode{
			Val:   1,
			Left:  &util.TreeNode{Val: 2},
			Right: &util.TreeNode{Val: 3},
		},
		expected: true,
	},
}

func TestIsSameTreeRecursion(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSameTree_recursion(tt.p, tt.q)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestIsSameTreeIteration(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSameTree_iteration(tt.p, tt.q)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
