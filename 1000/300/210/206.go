package leetcode_210

import "leetcode/util"

// Time complexity: O(n)
// Space complexity: O(1)
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

// Time complexity: O(n)
// Space complexity: O(n)
func reverseList_recursive(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList_recursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
