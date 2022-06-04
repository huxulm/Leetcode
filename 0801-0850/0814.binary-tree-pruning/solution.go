package binarytreepruning

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := pruneTree(root.Left)
	right := pruneTree(root.Right)

	if root.Val == 0 && left == right {
		return nil
	}
	root.Left, root.Right = left, right
	return root
}
