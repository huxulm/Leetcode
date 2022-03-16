package kthsmallestelementinabst

import (
	. "lc/structures"
)

func kthSmallest(root *TreeNode, k int) (ans int) {
	idx := 0
	var traversal func(root *TreeNode, k int)

	traversal = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		traversal(root.Left, k)
		// 中序遍历代码位置
		idx++
		// 找到第 k 小的元素
		if idx == k {
			ans = root.Val
		}
		traversal(root.Right, k)
	}

	traversal(root, k)
	return
}
