package flattenbinarytreetolinkedlist

import (
	. "lc/structures"
)

func flatten(root *TreeNode) {
	flattenHelper(root)
}

func flattenHelper(root *TreeNode) (head, tail *TreeNode) {
	if root == nil {
		return nil, nil
	}
	head = root
	if root.Left == nil && root.Right == nil {
		tail = head
	} else if root.Left == nil && root.Right != nil {
		_, tail = flattenHelper(root.Right)
	} else if root.Left != nil && root.Right == nil {
		root.Right, tail = flattenHelper(root.Left)
		root.Left = nil
	} else if root.Left != nil && root.Right != nil {
		leftHead, leftTail := flattenHelper(root.Left)
		rightHead, rightTail := flattenHelper(root.Right)
		leftTail.Right = rightHead
		head.Right = leftHead
		head.Left = nil
		tail = rightTail
		tail.Left = nil
	}
	return head, tail
}
