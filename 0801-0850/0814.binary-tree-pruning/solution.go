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

func pruneTree1(root *TreeNode) *TreeNode {
	if containsOne(root) {
		return root
	} else {
		return nil
	}
}

func containsOne(root *TreeNode) bool {
	if root == nil {
		return false
	}

	l, r := containsOne(root.Left), containsOne(root.Right)
	if !l {
		root.Left = nil
	}
	if !r {
		root.Right = nil
	}
	return root.Val == 1 || l || r
}
