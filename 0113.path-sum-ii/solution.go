package pathsumii

import (
	. "lc/structures"
)

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {

	track := []int{}

	var dfs func(root *TreeNode, targetSum int, track *[]int)
	dfs = func(root *TreeNode, targetSum int, track *[]int) {

		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			if targetSum == root.Val {
				*track = append(*track, root.Val)
				ans = append(ans, append([]int(nil), *track...))
				*track = (*track)[:len(*track)-1]
			}
			return
		}

		// left
		if root.Left != nil {
			*track = append(*track, root.Val)
			dfs(root.Left, targetSum-root.Val, track)
			*track = (*track)[:len(*track)-1]
		}

		if root.Right != nil {
			// right
			*track = append(*track, root.Val)
			dfs(root.Right, targetSum-root.Val, track)
			*track = (*track)[:len(*track)-1]
		}

	}

	dfs(root, targetSum, &track)

	return ans
}
