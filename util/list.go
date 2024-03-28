package util

type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateLinkedList Utility function to create a linked list from a slice.
func CreateLinkedList(elements []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, element := range elements {
		current.Next = &ListNode{
			Val:  element,
			Next: nil,
		}
		current = current.Next
	}
	return dummy.Next
}

// LinkedListToSlice Utility function to convert a linked list back to a slice.
func LinkedListToSlice(head *ListNode) []int {
	var elements []int

	for current := head; current.Next != nil; current = current.Next {
		elements = append(elements, current.Val)
	}

	return elements
}
