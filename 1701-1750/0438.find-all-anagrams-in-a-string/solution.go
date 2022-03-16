package findallanagramsinastring

// 输入: s = "cbaebabacd", p = "abc"
// 输入: s = "abab", p = "ab"
func findAnagrams(s string, p string) (ans []int) {
	n, m := len(s), len(p)
	if n < m {
		return
	}
	window, pm := [26]int{}, [26]int{}
	// 统计p中每个字符的个数
	for i := range p {
		pm[p[i]-'a']++
	}
	for i := range s[:m-1] {
		window[s[i]-'a']++
	}
	for i := m - 1; i < n; i++ {
		window[s[i]-'a']++
		if checkEqual(window, pm) {
			// if window == pm {
			ans = append(ans, i-m+1)
		}
		window[s[i-m+1]-'a']--
	}
	return
}

func checkEqual(win, m [26]int) bool {
	for i := range win {
		if win[i] != m[i] {
			return false
		}
	}
	return true
}
