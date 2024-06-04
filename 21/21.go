package merge_two_sorted_lists

import "leetcode/util"

// Time complexity: O(m+n)
// Space complexity: O(1)
func mergeTwoLists_iteration(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{}
	current := dummy

	p1, p2 := list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			current.Next = p2
			p2 = p2.Next
		} else {
			current.Next = p1
			p1 = p1.Next
		}

		current = current.Next
	}

	if p1 != nil {
		current.Next = p1
	} else {
		current.Next = p2
	}

	return dummy.Next
}

// Time complexity: O(m+n)
// Space complexity: O(m+n)
func mergeTwoLists_recursion(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists_recursion(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists_recursion(list1, list2.Next)
		return list2
	}
}
