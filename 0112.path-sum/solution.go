package pathsum

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 回溯
func hasPathSum(root *TreeNode, targetSum int) bool {
	return backtrack(root, &targetSum)
}

func backtrack(root *TreeNode, track *int) bool {
	if root == nil {
		return *track == 0
	}

	if root.Left == nil && root.Right == nil {
		return *track == root.Val
	}

	// 选择
	// 走左
	*track -= root.Val
	left := backtrack(root.Left, track)
	*track += root.Val
	if left {
		return true
	}
	// 走右侧
	*track -= root.Val
	right := backtrack(root.Right, track)
	*track += root.Val
	if right {
		return true
	}

	return false
}

// 动态规划
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == root.Right && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
}
