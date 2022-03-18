package deleteoperationfortwostrings

// 方法一: 最长公共子序列（LCS）
func minDistance(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return m + n - dp[m][n]*2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 方法二：动态规划
func minDistance1(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}

			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}

	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
