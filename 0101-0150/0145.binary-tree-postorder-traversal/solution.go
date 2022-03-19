package binarytreepostordertraversal

import (
	. "lc/structures"
)

func postorderTraversal(root *TreeNode) (ans []int) {
	traversal(&ans, root)
	return
}

func traversal(arr *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traversal(arr, root.Left)
	}
	if root.Right != nil {
		traversal(arr, root.Right)
	}
	*arr = append(*arr, root.Val)
}
