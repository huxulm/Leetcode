package reverselinkedlistii

import (
	. "lc/structures"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	// 保存在开始反转处节点的前一节点(可能为空,断裂处的位置)
	dummy := &ListNode{Next: head}
	cur := dummy

	// 保存被反转部分链表的头和尾
	var reverseHead, reverseTail *ListNode = nil, nil

	for i, n := 1, head; n != nil; i++ {
		next := n.Next
		// cur在 i==left 时断裂
		if i < left {
			cur = cur.Next
		}

		if i == left {
			reverseTail = n
			reverseTail.Next = nil
		}

		// 反转 [left,right]
		if i >= left && i <= right {
			n.Next = reverseHead
			reverseHead = n
		}

		// 合并三个部分的链表(前一部分>>反转部分>>后一部分)
		if i == right {
			cur.Next = reverseHead
			reverseTail.Next = next
			break
		}
		n = next
	}
	return dummy.Next
}
