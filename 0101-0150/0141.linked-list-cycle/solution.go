package linkedlistcycle

import (
	. "lc/structures"
)

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	// fast到达末尾停止
	for fast != nil && fast.Next != nil {
		// slow前进一步
		slow = slow.Next
		// fast前进两步
		fast = fast.Next.Next
		// 相遇
		if slow == fast {
			return true
		}
	}
	return false
}
