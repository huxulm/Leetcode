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

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// 找到倒数第n+1个节点
	x := findNthFromEnd(dummy, n+1)
	x.Next = x.Next.Next
	return dummy.Next
}

func findNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	// p1先走n步
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
