package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var slow, fast *ListNode = head, head.Next

	for ; fast != nil; fast = fast.Next {
		if fast.Val != slow.Val { // 重复
			slow.Next = fast
			slow = slow.Next
		}
		if fast.Next == nil {
			slow.Next = nil
		}
	}
	return head
}
