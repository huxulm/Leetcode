package swapnodesinpairs

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := first.Next
	head = second
	var prev *ListNode

	for first != nil && second != nil {
		first.Next = second.Next
		second.Next = first
		if prev != nil {
			prev.Next = second
		}
		prev = first
		first = first.Next
		if first != nil {
			second = first.Next
		}
	}

	return head
}
