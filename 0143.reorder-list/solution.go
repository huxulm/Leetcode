package reorderlist

import (
	"container/list"
	. "lc/structures"
)

func reorderList(head *ListNode) {
	dq := list.New()

	for n := head; n != nil; n = n.Next {
		dq.PushBack(n)
	}

	var cur = &ListNode{}

	for dq.Len() > 0 {
		h := dq.Remove(dq.Front()).(*ListNode)
		h.Next = nil
		cur.Next = h
		cur = cur.Next
		if dq.Len() > 0 {
			t := dq.Remove(dq.Back()).(*ListNode)
			t.Next = nil
			cur.Next = t
			cur = cur.Next
		}
	}
}

// 找中点+反转+合并
func reorderList1(head *ListNode) {

	var reverse func(head *ListNode) *ListNode

	reverse = func(head *ListNode) *ListNode {
		dummy := &ListNode{}
		cur := head
		for cur != nil {
			next := cur.Next
			cur.Next = dummy.Next
			dummy.Next = cur
			cur = next
		}
		return dummy.Next
	}

	var middleNode func(head *ListNode) *ListNode

	middleNode = func(head *ListNode) *ListNode {
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			// 到达末尾
			if fast == nil {
				break
			}
		}
		return slow
	}

	var mergeList func(list1, list2 *ListNode) *ListNode

	mergeList = func(list1, list2 *ListNode) *ListNode {
		dummy := &ListNode{}
		cur := dummy

		var temp1, temp2 *ListNode
		for list2 != nil {
			temp1, temp2 = list1.Next, list2.Next
			list2.Next = nil
			list1.Next = list2
			cur.Next = list1
			cur = list2
			list1, list2 = temp1, temp2
		}
		cur.Next = list1
		if list1 != nil && list1.Next != nil {
			cur.Next.Next = nil
		}
		return dummy.Next
	}

	if head == nil {
		return
	}

	mergeList(head, reverse(middleNode(head)))
}
