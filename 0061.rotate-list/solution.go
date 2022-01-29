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
