package uniquesubstringsinwraparoundstring

// 方法一: 双指针(也可以简化, 见方法二)
func findSubstringInWraproundString(p string) (ans int) {

	n := len(p)
	memo := [26]int{} // 记录以字母 p[i] 开头的最长子串长度

	// p[i...j]
	// 搜索p[i]开头的最长子串
	l, r := 0, 0
	for l < n {
		for r < n-1 && (p[r+1]-'a') == (p[r]-'a'+1)%26 {
			r++
		}
		// 更新 p[l]...p[r] 开头的最长长度
		for i := l; i <= r; i++ {
			if d := i - l + 1; d > memo[p[i]-'a'] { // 第i(i>=1)次遇到p[i]开头的子串
				ans += d - memo[p[i]-'a'] // 新增的子串个数
				memo[p[i]-'a'] = d        // 更新p[i]开头最长子串长度
			}
		}
		r++
		l = r
	}

	return
}

// 动态规划
func findSubstringInWraproundString1(p string) (ans int) {
	dp := [26]int{}
	k := 0
	for i := range p {
		if i > 0 && p[i]-'a' == (p[i-1]-'a'+1)%26 {
			k++
		} else {
			k = 1
		}
		dp[p[i]-'a'] = max(dp[p[i]-'a'], k)
	}
	for _, v := range dp {
		ans += v
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
