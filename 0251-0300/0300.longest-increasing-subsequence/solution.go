package longestincreasingsubsequence

import (
	"sort"
)

// 方法二: 动态规划
func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	// dp[i] s[i]结尾最长递增子序列长度
	dp := make([]int, n)
	dp[0] = 1

	ans = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 方法二: 动态规划(使用最长公共子序列方法求解)
func lengthOfLIS1(nums []int) (ans int) {
	m := len(nums)
	arr := make([]int, m)
	copy(arr, nums)
	sort.Ints(arr)

	// 去除重复元素
	i, j := 0, 0
	for j < m {
		if arr[i] != arr[j] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
		j++
	}

	arr = arr[:i+1]
	n := len(arr)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums[i-1] == arr[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// 方法三：贪心 + 二分查找
