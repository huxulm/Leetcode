package longesthappyprefix

func longestPrefix(s string) (ans string) {
	n := len(s)
	if n <= 1 {
		return
	}
	// 初始假设最长情况
	k := n - 1
	for k > 0 && s[:k] != s[n-k:n] {
		k--
	}
	if k > 0 {
		ans = s[:k]
	}
	return
}
