package removenthnodefromendoflist

import "container/list"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	stack := list.New()
	for n := head; n != nil; n = n.Next {
		stack.PushFront(n)
	}

	var removed *ListNode
	for n > 0 {
		removed = stack.Remove(stack.Front()).(*ListNode)
		n--
	}

	if stack.Len() == 0 {
		return removed.Next
	}

	// remove一定不为nil
	head = stack.Front().Value.(*ListNode)
	head.Next = removed.Next
	return head
}
