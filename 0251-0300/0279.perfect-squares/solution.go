package perfectsquares

func numSquares(n int) int {
	// 我们可以依据题目的要求写出状态表达式：dp[i]表示最少需要多少个数的平方来表示整数 i。
	// i 的完全平方数必然在 [1, sqrt(n)]之间

	// dp[i] 第i个数的最少完全平方数
	// dp[i] = 1 + min{dp[i-j*j]} (1<=j<=sqrt(i))
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		minn := 1<<31 - 1
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}

	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
