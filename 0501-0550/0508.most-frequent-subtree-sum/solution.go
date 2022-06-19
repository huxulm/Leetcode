package mostfrequentsubtreesum

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func findFrequentTreeSum(root *TreeNode) (ans []int) {
	mp := map[int]int{}
	// sum of tree
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) (s int) {
		if root == nil {
			return 0
		}
		ls := dfs(root.Left)
		rs := dfs(root.Right)
		s = ls + rs + root.Val
		mp[s]++
		return s
	}
	dfs(root)
	maxCnt := 0
	for k, v := range mp {
		if v > maxCnt {
			ans = []int{k}
			maxCnt = v
		} else if v == maxCnt {
			ans = append(ans, k)
		}
	}
	return
}
