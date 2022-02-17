package deletenodeinalinkedlist

import (
	. "lc/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	prev := node
	for cur := node.Next; cur != nil; cur = cur.Next {
		prev.Val = cur.Val
		if cur.Next != nil {
			prev = cur
		}
	}
	prev.Next = nil
}
