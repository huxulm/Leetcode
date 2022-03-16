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

		*track = append(*track, root.Val)
		targetSum -= root.Val

		defer func() {
			*track = (*track)[:len(*track)-1]
		}()

		if root.Left == nil && root.Right == nil && targetSum == 0 {
			ans = append(ans, append([]int(nil), *track...))
			return
		}

		dfs(root.Left, targetSum-root.Val, track)
		dfs(root.Right, targetSum-root.Val, track)
	}

	dfs(root, targetSum, &track)

	return ans
}
