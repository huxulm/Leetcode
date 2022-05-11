package integerbreak

import "math"

// 方法一: 答案对 但 方法不正确
func integerBreak(n int) (ans int) {
	dp := n / 2 * (n/2 + n%2)
	ans = dp
	for i := 2; i < n; i++ {
		dp = max(int(math.Pow(float64(n/i), float64(i-1)))*(n/i+n%i), int(math.Pow(float64(n/i+1), float64(i-1)))*(n-(i-1)*(n/i+1)))
		if dp > ans {
			ans = dp
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

// 动态规划
func integerBreak1(n int) (ans int) {
	// dp[i] n 拆分成 i 个整数乘积的最大值
	dp := make([]int, n+1)

	dp[2] = n / 2 * (n/2 + n%2)

	return
}

// dp[i] 拆分数字i的最大乘积
func integerBreak2(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}
