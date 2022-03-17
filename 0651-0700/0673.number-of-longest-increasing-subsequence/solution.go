package numberoflongestincreasingsubsequence

func findNumberOfLIS(nums []int) (ans int) {
	n := len(nums)
	// dp[i] 表示以 nums[i] 结尾的最长上升子序列的长度
	f := make([]int, n)
	// cnt[i] 表示以 nums[i] 结尾的最长上升子序列的个数
	cnt := make([]int, n)
	maxLen := 0

	for i, x := range nums {
		f[i] = 1
		cnt[i] = 1
		for j, y := range nums[:i] {
			if x > y {
				if f[j]+1 > f[i] {
					f[i] = f[j] + 1
					// 重新计数
					cnt[i] = cnt[j]
				} else if f[j]+1 == f[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if f[i] > maxLen {
			ans = cnt[i] // 重置计数
			maxLen = f[i]
		} else if f[i] == maxLen {
			ans += cnt[i]
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
