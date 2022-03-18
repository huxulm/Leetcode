package longestcommonsubsequence

func longestCommonSubsequence(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	// dp[i][j] 定义: s1[:i] 与 s2[:j] 的 LCS
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				// dp[i][j] = 0
				continue
			}
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
