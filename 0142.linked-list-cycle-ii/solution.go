package linkedlistcycleii

import (
	. "lc/structures"
)

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	// 重新指向头节点
	slow = head
	// 快慢指针同步前进，相交点就是环起点
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}
