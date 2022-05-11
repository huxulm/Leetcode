package invertbinarytree

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 迭代
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sta := []*TreeNode{root}
	for len(sta) > 0 {
		node := sta[0]
		sta = sta[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			sta = append(sta, node.Left)
		}
		if node.Right != nil {
			sta = append(sta, node.Right)
		}
	}
	return root
}
