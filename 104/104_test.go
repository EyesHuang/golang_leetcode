package maximum_depth_of_binary_tree

import "testing"

var tests = []struct {
	name string
	root *TreeNode
	want int
}{
	{"Test with an empty tree", nil, 0},
	{"Test with a single node", &TreeNode{Val: 1}, 1},
	{
		"Test with a balanced tree",
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 6},
			},
		},
		3,
	},
	{
		"Test with a skewed tree (left)",
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 4},
				},
			},
		},
		4,
	},
	{
		"Test with a skewed tree (right)",
		&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
		},
		4,
	},
}

func TestXxx(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.root)
			if got != tt.want {
				t.Errorf("maxDepth(%v) = %d; want %d", tt.root, got, tt.want)
			}
		})
	}
}
