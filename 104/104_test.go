package maximum_depth_of_binary_tree

import (
	"testing"

	"leetcode/util"
)

var tests = []struct {
	name string
	root *util.TreeNode
	want int
}{
	{"Test with an empty tree", nil, 0},
	{"Test with a single node", &util.TreeNode{Val: 1}, 1},
	{
		"Test with a balanced tree",
		&util.TreeNode{
			Val: 1,
			Left: &util.TreeNode{
				Val:   2,
				Left:  &util.TreeNode{Val: 4},
				Right: &util.TreeNode{Val: 5},
			},
			Right: &util.TreeNode{
				Val:   3,
				Right: &util.TreeNode{Val: 6},
			},
		},
		3,
	},
	{
		"Test with a skewed tree (left)",
		&util.TreeNode{
			Val: 1,
			Left: &util.TreeNode{
				Val: 2,
				Left: &util.TreeNode{
					Val:  3,
					Left: &util.TreeNode{Val: 4},
				},
			},
		},
		4,
	},
	{
		"Test with a skewed tree (right)",
		&util.TreeNode{
			Val: 1,
			Right: &util.TreeNode{
				Val: 2,
				Right: &util.TreeNode{
					Val:   3,
					Right: &util.TreeNode{Val: 4},
				},
			},
		},
		4,
	},
}

func TestMaxDepth(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.root)
			if got != tt.want {
				t.Errorf("maxDepth(%v) = %d; want %d", tt.root, got, tt.want)
			}
		})
	}
}
