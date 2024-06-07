package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time complexity: O(n)
// Space complexity: O(n)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
