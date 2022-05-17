package binarytreeinordertraversal

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 回溯算法思路
func inorderTraversal(root *TreeNode) []int {
	var Traversal func(root *TreeNode, arr *[]int)

	Traversal = func(root *TreeNode, arr *[]int) {
		if root == nil {
			return
		}
		Traversal(root.Left, arr)
		*arr = append(*arr, root.Val)
		Traversal(root.Right, arr)
	}

	arr := make([]int, 0)
	Traversal(root, &arr)
	return arr
}

// 动态规划思路
func inorderTraversal1(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}
	arr = append(arr, inorderTraversal1(root.Left)...)
	arr = append(arr, root.Val)
	arr = append(arr, inorderTraversal1(root.Right)...)
	return arr
}

// 迭代法
func inorderTraversal2(root *TreeNode) (ans []int) {
	sta := []*TreeNode{}
	cur := root
	for cur != nil || len(sta) > 0 {
		// 左侧优先压栈
		for cur != nil {
			sta = append(sta, cur)
			cur = cur.Left
		}
		// pop栈顶元素
		cur = sta[len(sta)-1]
		ans = append(ans, cur.Val)
		sta = sta[:len(sta)-1]
		cur = cur.Right
	}
	return
}
