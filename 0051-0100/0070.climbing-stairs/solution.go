// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 注意：给定 n 是一个正整数。
package climbingstairs

// 动态规划: 自底向上
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	// 最后一次爬1阶或爬2阶
	// dp[i] = dp[i-1] + dp[i-2]
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 动态规划: 自底向上
func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	// 最后一次爬1阶或爬2阶
	// dp[i] = dp[i-1] + dp[i-2]
	dp1, dp2 := 1, 2
	tmp := 0
	for i := 3; i <= n; i++ {
		tmp = dp1 + dp2
		dp1, dp2 = dp2, tmp
	}
	return dp2
}
