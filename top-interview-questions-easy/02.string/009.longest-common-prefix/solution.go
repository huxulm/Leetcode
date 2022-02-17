package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	ans := ""
	preffix := strs[0]
	for i := 1; i < len(strs); i++ {
		preffix = commonPrefix(preffix, strs[i])
		if len(preffix) > len(ans) {
			ans = preffix
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func commonPrefix(s1, s2 string) string {
	n1, n2 := len(s1), len(s2)
	idx := 0
	for i := 0; i < min(n1, n2); i++ {
		if s1[i] == s2[i] {
			idx++
		} else {
			break
		}
	}
	return s1[:idx+1]
}
