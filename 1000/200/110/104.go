package leetcode_110

import "leetcode/util"

// Time complexity: O(n)
// Space complexity: O(n)
func maxDepth(root *util.TreeNode) int {
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
