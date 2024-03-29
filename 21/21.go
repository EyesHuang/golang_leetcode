package merge_two_sorted_lists

import "leetcode/util"

func mergeTwoLists(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {
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
