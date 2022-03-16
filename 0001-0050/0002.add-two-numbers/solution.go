package addtwonumbers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Next: &ListNode{}}
	cur := res
	for l1 != nil || l2 != nil {
		sum := cur.Next.Val
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		next := &ListNode{Val: sum}
		next.Next = &ListNode{Val: next.Val / 10}
		next.Val %= 10
		cur.Next = next
		cur = next
	}
	if cur.Next.Val == 0 {
		cur.Next = nil
	}
	return res.Next
}
