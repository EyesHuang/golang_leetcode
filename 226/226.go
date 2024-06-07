package invert_binary_tree

import "leetcode/util"

// Time complexity: O(n)
// Space complexity: O(n)
func invertTree_recursion(root *util.TreeNode) *util.TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree_recursion(root.Left)
	right := invertTree_recursion(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// Time complexity: O(n)
// Space complexity: O(n)
func invertTree_iteration(root *util.TreeNode) *util.TreeNode {
	if root == nil {
		return nil
	}

	queue := []*util.TreeNode{root}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		tmp := current.Left
		current.Left = current.Right
		current.Right = tmp

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return root
}
