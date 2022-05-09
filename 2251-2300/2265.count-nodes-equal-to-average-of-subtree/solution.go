package countnodesequaltoaverageofsubtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) (ans int) {
	var dfs func(root *TreeNode) (sum, cnt int)
	dfs = func(root *TreeNode) (sum int, cnt int) {
		if root == nil {
			return 0, 0
		}
		ls, lc := dfs(root.Left)
		rs, rc := dfs(root.Right)
		cnt = 1 + lc + rc
		sum = root.Val + ls + rs
		if sum/cnt == root.Val {
			ans++
		}
		return
	}
	dfs(root)
	return
}
