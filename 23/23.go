package merge_k_sorted_lists

import (
	"leetcode/util"
	"sort"
)

func mergeKLists_BruteForce(lists []*util.ListNode) *util.ListNode {
	nodes := make([]int, 0, 10000)

	for _, l := range lists {
		for l != nil {
			nodes = append(nodes, l.Val)
			l = l.Next
		}
	}

	sort.Ints(nodes)

	dummy := &util.ListNode{}
	currNode := dummy

	for _, node := range nodes {
		currNode.Next = &util.ListNode{Val: node}
		currNode = currNode.Next
	}

	return dummy.Next
}
