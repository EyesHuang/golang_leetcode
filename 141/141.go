package linked_list_cycle

import "leetcode/util"

// Time complexity: O(n)
// Space complexity: O(n)
func hasCycle_set(head *util.ListNode) bool {
	type void struct{}
	var member void
	set := make(map[*util.ListNode]void)
	current := head

	for current != nil {
		if _, ok := set[current]; ok {
			return true
		}

		set[current] = member
		current = current.Next
	}

	return false
}

// Time complexity: O(n)
// Space complexity: O(1)
func hasCycle_twoPointers(head *util.ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
