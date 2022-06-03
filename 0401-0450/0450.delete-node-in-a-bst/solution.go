package deletenodeinabst

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}
		t := root
		root = min(t.Right)
		root.Right = deleteMin(t.Right)
		root.Left = t.Left
	}
	return root
}

func deleteMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root.Right
	}
	root.Left = deleteMin(root.Left)
	return root
}

func min(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return min(root.Left)
}

// 递归
func deleteNode1(root *TreeNode, key int) *TreeNode {
	// 设 x 为要删除的节点
	// 找到 x 的后驱节点（x右子树最小）
	// successor = min(x.right)
	// successor.left = x.left
	// successor.right = delete(x.right, successor.val)
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode1(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode1(root.Right, key)
		return root
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode1(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
}
