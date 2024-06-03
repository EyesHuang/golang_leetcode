package reverse_linked_list

import "leetcode/util"

func reverseList_iterative(head *util.ListNode) *util.ListNode {
	var prev *util.ListNode
	curr := head

	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}

	return prev
}
