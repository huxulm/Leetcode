package interleavingstring

// 参考: 不同路径 https://leetcode-cn.com/problems/unique-paths/
// 思路: 动态规划
func isInterleave(s1 string, s2 string, s3 string) (ans bool) {
	n1, n2, n3 := len(s1), len(s2), len(s3)

	if n1+n2 != n3 {
		return false
	}

	// dp[i][j] =
	dp := make([][]bool, n1+1)

	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}

	dp[0][0] = true

	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[p]
			}
		}
	}

	return dp[n1][n2]
}

func isInterleave1(s1 string, s2 string, s3 string) (ans bool) {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return
	}

	// dp[i][j] 表示 s1[:i] 和 s[:j] 能否交错构成 s3[:i+j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s3[p] == s1[i-1]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s3[p] == s2[j-1])
			}
		}
	}
	return dp[m][n]
}
