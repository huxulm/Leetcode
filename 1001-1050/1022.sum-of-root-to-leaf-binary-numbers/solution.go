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

// 二叉树迭代遍历
func sumRootToLeaf2(root *TreeNode) (ans int) {
	val, sta := 0, []*TreeNode{}
	cur := root
	var pre *TreeNode
	for cur != nil || len(sta) > 0 {
		for cur != nil {
			sta = append(sta, cur)
			val = val<<1 | cur.Val
			cur = cur.Left
		}

		cur = sta[len(sta)-1]

		if cur.Right == nil || cur.Right == pre {
			// 当前为叶子节点
			if cur.Left == nil && cur.Right == nil {
				ans += val
			}
			val >>= 1
			sta = sta[:len(sta)-1]
			pre = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return
}

// BFS
func sumRootToLeaf3(root *TreeNode) (ans int) {
	q := []*TreeNode{root}
	for q != nil {
		tmp := q
		q = nil
		for _, n := range tmp {
			if n.Left != nil {
				n.Left.Val = n.Val<<1 | n.Left.Val
				q = append(q, n.Left)
			}
			if n.Right != nil {
				n.Right.Val = n.Val<<1 | n.Right.Val
				q = append(q, n.Right)
			}
			if n.Left == n.Right {
				ans += n.Val
			}
		}
	}
	return
}
