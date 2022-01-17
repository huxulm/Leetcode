package implementstrstr

// 暴力, O(n*m)
func strStr(s1 string, s2 string) int {

	l1, l2 := len(s1), len(s2)

	if l2 == 0 {
		return 0
	}

	for i := 0; i < l1; i++ {
		if s1[i] == s2[0] {
			// 检查 s1[i...i+l2] 是否和s2相等
			if l1-i < l2 {
				return -1
			}
			if s1[i:i+l2] == s2 {
				return i
			}
		}
	}

	return -1
}

// 有限状态机之 KMP 字符匹配算法
func strStr1(s1 string, s2 string) int {
	dfa := KMP(s2)

	N, M := len(s1), len(s2)
	var i, j = 0, 0
	for ; i < N && j < M; i++ {
		j = dfa[s1[i]][j]
		if j == M {
			return i - M + 1
		}
	}
	return -1
}

func KMP(pat string) [][]int {
	M := len(pat)
	R := 256

	dfa := make([][]int, R)

	for i := range dfa {
		dfa[i] = make([]int, M)
	}

	dfa[pat[0]][0] = 1
	for X, j := 0, 1; j < M; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][X]
		}
		dfa[pat[j]][j] = j + 1
		X = dfa[pat[j]][X]
	}
	return dfa
}

// 一维状态数组的KMP
func strStr2(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
