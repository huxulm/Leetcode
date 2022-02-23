package constructbinarytreefrompreorderandinordertraversal

import (
	. "lc/structures"
)

// 前序+中序构建二叉树
// 思路:
// 	 1. 前序遍历找到根节点
//   2. 搜索根节点在中序号遍历的位置，将中序遍历分割成左右子树的中序遍历，
// 			同时获取左右子树的大小
//   3. 通过左右子树的大小，将前序遍历也分成左右子树的前序遍历
//   4. 递归调用，构建左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	rootVal := preorder[0]
	rootIndex := -1

	for i := range inorder {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}

	root := &TreeNode{Val: rootVal}

	// 左右子树的大小
	leftSize, rightSize := rootIndex, len(preorder)-rootIndex-1

	if leftSize > 0 {
		root.Left = buildTree(preorder[1:leftSize+1], inorder[:leftSize])
	}
	if rightSize > 0 {
		root.Right = buildTree(preorder[leftSize+1:], inorder[leftSize+1:])
	}
	return root
}
