package reversenodesinkgroup

import (
	. "lc/structures"
)

func reverseKGroup(head *ListNode, k int) *ListNode {

	if head == nil || k < 2 {
		return head
	}

	var dummy *ListNode = &ListNode{Next: head}

	return dummy.Next
}
