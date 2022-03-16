// https://leetcode-cn.com/problems/balanced-binary-tree/
package balancedbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// 一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
	if root == nil {
		return true
	}
	return abs(depth(root.Left)-depth(root.Right)) > 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isBalanced1(root *TreeNode) bool {
	// 一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
	if root == nil {
		return true
	}
	isBalanced := true
	depth1(root, &isBalanced)
	return isBalanced
}

func depth1(root *TreeNode, isBlanced *bool) int {
	if root == nil {
		return 0
	}
	// 结束递归
	if !*isBlanced {
		return -1
	}

	leftDepth := depth1(root.Left, isBlanced)
	rightDepth := depth1(root.Right, isBlanced)
	// 后序遍历位置
	if abs(leftDepth-rightDepth) > 1 {
		*isBlanced = false
	}
	return 1 + max(leftDepth, rightDepth)
}
