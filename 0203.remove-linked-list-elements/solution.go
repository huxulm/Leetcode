package removelinkedlistelements

import (
	. "lc/structures"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	for prev := dummy; prev.Next != nil; {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return dummy.Next
}

// LC official
func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements1(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
