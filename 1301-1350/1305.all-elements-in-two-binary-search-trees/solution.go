package allelementsintwobinarysearchtrees

import (
	. "lc/structures"
)

// 方法一: 中序 + 归并
func getAllElements(root1, root2 *TreeNode) (ans []int) {
	var inorder func(root *TreeNode) (res []int)
	inorder = func(root *TreeNode) (res []int) {
		var dfs func(root *TreeNode)
		dfs = func(root *TreeNode) {
			if root == nil {
				return
			}
			dfs(root.Left)
			res = append(res, root.Val)
			dfs(root.Right)
		}
		dfs(root)
		return
	}
	arr1, arr2 := inorder(root1), inorder(root2)
	m, n := len(arr1), len(arr2)
	merged := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m && p2 == n {
			return merged
		}
		if p2 == n || (p1 < m && arr1[p1] < arr2[p2]) {
			merged = append(merged, arr1[p1])
			p1++
		} else if p1 == m || (p2 < n && arr1[p1] >= arr2[p2]) {
			merged = append(merged, arr2[p2])
			p2++
		}
	}
}
