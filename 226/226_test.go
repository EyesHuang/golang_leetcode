package invert_binary_tree

import (
	"reflect"
	"testing"

	"leetcode/util"
)

var tests = []struct {
	name     string
	input    *util.TreeNode
	expected *util.TreeNode
}{
	{
		name:     "NilTree",
		input:    nil,
		expected: nil,
	},
	{
		name:     "SingleNode",
		input:    &util.TreeNode{Val: 1},
		expected: &util.TreeNode{Val: 1},
	},
	{
		name: "FullTree",
		input: &util.TreeNode{
			Val: 4,
			Left: &util.TreeNode{
				Val:   2,
				Left:  &util.TreeNode{Val: 1},
				Right: &util.TreeNode{Val: 3},
			},
			Right: &util.TreeNode{
				Val:   7,
				Left:  &util.TreeNode{Val: 6},
				Right: &util.TreeNode{Val: 9},
			},
		},
		expected: &util.TreeNode{
			Val: 4,
			Left: &util.TreeNode{
				Val:   7,
				Left:  &util.TreeNode{Val: 9},
				Right: &util.TreeNode{Val: 6},
			},
			Right: &util.TreeNode{
				Val:   2,
				Left:  &util.TreeNode{Val: 3},
				Right: &util.TreeNode{Val: 1},
			},
		},
	},
	{
		name: "LeftSkewedTree",
		input: &util.TreeNode{
			Val: 1,
			Left: &util.TreeNode{
				Val: 2,
				Left: &util.TreeNode{
					Val: 3,
				},
			},
		},
		expected: &util.TreeNode{
			Val: 1,
			Right: &util.TreeNode{
				Val: 2,
				Right: &util.TreeNode{
					Val: 3,
				},
			},
		},
	},
	{
		name: "RightSkewedTree",
		input: &util.TreeNode{
			Val: 1,
			Right: &util.TreeNode{
				Val: 2,
				Right: &util.TreeNode{
					Val: 3,
				},
			},
		},
		expected: &util.TreeNode{
			Val: 1,
			Left: &util.TreeNode{
				Val: 2,
				Left: &util.TreeNode{
					Val: 3,
				},
			},
		},
	},
}

func TestInvertTreeRecursion(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := invertTree_recursion(tt.input)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}

func TestInvertTreeIteration(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := invertTree_iteration(tt.input)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
