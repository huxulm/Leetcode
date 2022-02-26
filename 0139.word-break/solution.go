package wordbreak

// DFS + 剪枝
func wordBreak(s string, wordDict []string) bool {
	n := len(s)

	// 从s[start:]开始搜索词典
	// 前start个字符已被匹配过已被搜索
	var dfs func(start int, vis []bool)
	dfs = func(start int, vis []bool) {
		if start == n {
			return
		}
		for i := range wordDict {
			wlen := len(wordDict[i])
			// 剪枝: s的前面部分已经匹配成功的不再重新匹配
			if wlen <= n-start && wordDict[i] == s[start:wlen+start] && !vis[start+wlen-1] {
				vis[start+wlen-1] = true
				dfs(start+wlen, vis)
			}
		}
	}

	var vis = make([]bool, n)
	dfs(0, vis)

	return vis[n-1]
}

// 动态规划
func wordBreak1(s string, wordDict []string) bool {
	n := len(s)
	// dp[i] 表示字符串 s 前 i 个字符组成的字符串 s[0..i-1] 是否能被空格拆分成若干个字典中出现的单词
	dp := make([]bool, n+1)
	dp[0] = true

	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
