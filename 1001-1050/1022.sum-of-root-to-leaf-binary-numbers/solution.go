package sumofroottoleafbinarynumbers

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func sumRootToLeaf(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode, v int)
	dfs = func(root *TreeNode, v int) {
		if root == nil {
			return
		}
		v = v*2 + root.Val
		if root.Left == root.Right {
			ans += v
		}
		dfs(root.Left, v)
		dfs(root.Right, v)
	}
	dfs(root, 0)
	return
}

// LC 1022
func sumRootToLeaf1(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode, val int) int
	dfs = func(root *TreeNode, val int) int {
		if root == nil {
			return 0
		}
		val = val<<1 | root.Val
		if root.Left == root.Right {
			return val
		}
		return dfs(root.Left, val) + dfs(root.Right, val)
	}
	return dfs(root, 0)
}
