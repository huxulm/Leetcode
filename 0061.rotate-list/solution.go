package rotatelist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	size := 0

	for n := head; n != nil; n = n.Next {
		size++
	}

	if size < 2 || k%size == 0 {
		return head
	}

	k %= size

	var (
		leftHead, leftTrail, rightHead, rightTrail *ListNode
	)
	leftHead = head
	for i, n := 0, head; n != nil; {
		if size-1 == i+k {
			leftTrail = n
			rightHead = n.Next
			break
		}
		n, i = n.Next, i+1
	}

	for n := rightHead; n != nil; n = n.Next {
		if n.Next == nil {
			rightTrail = n
		}
	}

	leftTrail.Next = nil
	rightTrail.Next = leftHead
	return rightHead
}

// LC official
func rotateRight1(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
