package partition_list

import "leetcode/util"

func partition(head *util.ListNode, x int) *util.ListNode {
	dummy1 := &util.ListNode{}
	dummy2 := &util.ListNode{}

	p1, p2 := dummy1, dummy2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}

		tmp := head.Next
		head.Next = nil
		head = tmp
	}

	p1.Next = dummy2.Next

	return dummy1.Next
}
