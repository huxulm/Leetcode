package binarytreepreordertraversal

import (
	. "lc/structures"
)

func preorderTraversal(root *TreeNode) (ans []int) {
	traversal(&ans, root)
	return
}

func traversal(arr *[]int, root *TreeNode) {
	if root != nil {
		*arr = append(*arr, root.Val)
		traversal(arr, root.Left)
		traversal(arr, root.Right)
	}
}
