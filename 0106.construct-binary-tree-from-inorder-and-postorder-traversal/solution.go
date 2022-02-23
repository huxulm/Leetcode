package constructbinarytreefrominorderandpostordertraversal

import (
	. "lc/structures"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)

	if n == 0 {
		return nil
	}

	rootIndex := -1
	for i := range inorder {
		if inorder[i] == postorder[n-1] {
			rootIndex = i
			break
		}
	}

	// 构建根节点
	root := &TreeNode{Val: postorder[n-1]}

	if rootIndex >= 0 {
		// 构建左右子树
		root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
		root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:n-1])
	}

	return root
}
