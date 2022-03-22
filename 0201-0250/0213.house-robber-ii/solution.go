package houserobberii

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	memo1 := make([]int, n)
	memo2 := make([]int, n)
	for i := range memo1 {
		memo1[i] = -1
		memo2[i] = -1
	}

	// 两次调用使用两个不同的备忘录
	return max(dp(nums, 0, n-2, memo1), dp(nums, 1, n-1, memo2))
}

// 定义：计算闭区间 [start,end] 的最优结果
func dp(nums []int, start, end int, memo []int) int {
	if start > end {
		return 0
	}

	if memo[start] != -1 {
		return memo[start]
	}
	// 状态转移方程
	res := max(dp(nums, start+2, end, memo)+nums[start], dp(nums, start+1, end, memo))

	memo[start] = res
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob1(nums []int) int {
	return 0
}
