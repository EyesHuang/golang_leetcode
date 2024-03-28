package merge_two_sorted_lists

import "leetcode/util"

func mergeTwoLists(list1 *util.ListNode, list2 *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{
		Val:  -1,
		Next: nil,
	}

	p := dummy
	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}

		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}
