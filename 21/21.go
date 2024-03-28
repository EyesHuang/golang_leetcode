package merge_two_sorted_lists

import "leetcode/util"

func mergeTwoLists(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}
