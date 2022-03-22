package twosumivinputisabst

import (
	. "lc/structures"
)

func findTarget(root *TreeNode, k int) (ans bool) {

	var m = make(map[int]bool)

	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		if m[k-root.Val] {
			ans = true
			return
		}
		m[root.Val] = true
		traversal(root.Left)
		traversal(root.Right)
	}
	return
}

func findTarget1(root *TreeNode, k int) bool {

	var m = make(map[int]bool)

	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if m[k-root.Val] {
			return true
		}
		m[root.Val] = true
		return dfs(root.Left) || dfs(root.Right)
	}

	return dfs(root)
}
