package longestsubstringwithoutrepeatingcharacters

// 暴力穷举
func lengthOfLongestSubstring(s string) int {
	// dp[i] : s[0:i]最大无重复字串长度
	// 以s[i]结尾的最大子串长度
	res := 0
	for i := 0; i < len(s); i++ {
		len_i := 0
		rec := make(map[byte]struct{})
		for j := i; j >= 0; j-- {
			if _, ok := rec[s[j]]; ok {
				break
			} else {
				rec[s[j]] = struct{}{}
			}
			len_i++
		}
		// 更新答案
		res = max(res, len_i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	// 记录字符出现次数
	rec := make(map[byte]int)
	left, right := 0, 0
	n := len(s)
	res := 0
	for right < n {
		c := s[right]
		right++
		rec[c]++
		// 判断左侧窗口是否要收缩
		for rec[c] > 1 {
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			rec[d]--
		}
		res = max(res, right-left)
	}
	return res
}

// 滑动窗口
func lengthOfLongestSubstring2(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

// 滑动窗口
func lengthOfLongestSubstring3(s string) (ans int) {
	m := map[byte]int{}
	n := len(s)

	l, r := 0, 0
	for r < n {
		m[s[r]]++

		for l < r && m[s[r]] > 1 {
			m[s[l]]--
			l++
		}

		if r-l+1 > ans {
			ans = r - l + 1
		}
		r++
	}
	return
}
