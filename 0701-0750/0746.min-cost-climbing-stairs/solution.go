package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	// n >= 2
	n := len(cost)
	// 爬上第i层楼梯的最小花费
	dp := make([]int, n+1)
	dp[2] = min(cost[0], cost[1])

	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}

	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 空间降维
// F(N) = min(F(n-2)+Cost(n-2), F(N-1)+Cost(n-1))
func minCostClimbingStairs1(cost []int) int {
	// n >= 2
	n := len(cost)
	// 爬上第i层楼梯的最小花费
	dp1, dp2 := 0, min(cost[0], cost[1])

	for i := 3; i <= n; i++ {
		dp1, dp2 = dp2, min(dp1+cost[i-2], dp2+cost[i-1])
	}

	return dp2
}

// 动态规划
func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	if n < 2 {
		return 0
	}
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
