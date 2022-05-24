package univaluedbinarytree

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 迭代遍历
func isUnivalTree(root *TreeNode) bool {
	var v = -1

	sta := []*TreeNode{}
	cur := root
	for cur != nil || len(sta) > 0 {

		// 左侧压栈
		for cur != nil {
			sta = append(sta, cur)
			cur = cur.Left
		}

		// POP栈顶元素(二叉树最左)
		cur = sta[len(sta)-1]
		sta = sta[:len(sta)-1]

		// 中序遍历位置
		if v == -1 {
			v = cur.Val
		} else if v != cur.Val {
			return false
		}

		// 遍历当前节点右子树
		cur = cur.Right
	}
	return true
}
