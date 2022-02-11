package reverselinkedlistii

import (
	. "lc/structures"
)

// { 1, 2, 3, 4, 5, 6}
// [2,3]
// { 1, 3, 2, 4, 5, 6}
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	var reverseHead, reverseTail *ListNode = &ListNode{}, nil
	cur := dummy
	for i, n := 1, head; n != nil; i++ {
		next := n.Next
		if i < left {
			cur = cur.Next
		}
		if i == left {
			reverseTail = n
			reverseTail.Next = nil
		}
		if i >= left && i <= right {
			n.Next = reverseHead.Next
			reverseHead.Next = n
		}
		// 合并
		if i == right {
			cur.Next = reverseHead.Next
			reverseTail.Next = next
			break
		}
		n = next
	}
	return dummy.Next
}
