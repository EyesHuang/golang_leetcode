package same_tree

import "leetcode/util"

// Time complexity: O(n)
// Space complexity: O(n)
func isSameTree_recursion(p *util.TreeNode, q *util.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree_recursion(p.Left, q.Left) && isSameTree_recursion(p.Right, q.Right)
}

// Time complexity: O(n)
// Space complexity: O(n)
func isSameTree_iteration(p *util.TreeNode, q *util.TreeNode) bool {
	type NodePair struct {
		t1 *util.TreeNode
		t2 *util.TreeNode
	}

	queue := []NodePair{{p, q}}

	for len(queue) != 0 {
		pair := queue[0]
		queue = queue[1:]

		if !check(pair.t1, pair.t2) {
			return false
		}

		if pair.t1 != nil {
			queue = append(queue, NodePair{pair.t1.Left, pair.t2.Left}, NodePair{pair.t1.Right, pair.t2.Right})
		}
	}

	return true
}

func check(p, q *util.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val
}
