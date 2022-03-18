package reverselinkedlist

import (
	. "lc/structures"
)

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for n := head; n != nil; {
		next := n.Next
		n.Next = dummy.Next
		dummy.Next = n
		n = next
	}
	return dummy.Next
}
