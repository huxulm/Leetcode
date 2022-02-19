package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) (ans int) {
	n := len(s)
	m := make(map[byte]int)

	l, r := 0, 0

	for r < n {
		// 记录滑入窗口的字符
		m[s[r]]++
		// 判断窗口左侧是否需要收缩
		// 条件: 有重复字符收缩左侧窗口
		for l <= r && m[s[r]] > 1 {
			m[s[l]]--
			l++
		}

		// 用新的更大的窗口长度更新答案
		if r-l+1 > ans {
			ans = r - l + 1
		}
		// 窗口向右侧扩展
		r++
	}

	return
}
