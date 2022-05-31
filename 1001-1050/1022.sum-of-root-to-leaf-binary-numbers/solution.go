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
