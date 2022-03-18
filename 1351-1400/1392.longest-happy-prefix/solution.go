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

// KMP
func longestPrefix1(s string) (ans string) {
	n := len(s)
	next := make([]int, n)
	for j, i := 0, 1; i < n; i++ {
		for j > 0 && s[j] != s[i] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
	ans = s[:next[n-1]]
	return
}
