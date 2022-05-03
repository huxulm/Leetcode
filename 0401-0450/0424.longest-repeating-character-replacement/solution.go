package longestrepeatingcharacterreplacement

// 滑动窗口
// 思路:
// 我们可以枚举字符串中的每一个位置作为右端点，
// 然后找到其最远的左端点的位置，满足该区间内除了出现次数最多的那一类字符之外，
// 剩余的字符（即非最长重复字符）数量不超过 k 个。
// func characterReplacement(s string, k int) (ans int) {
// 	cnt := [26]int{}

// 	maxCnt, left := 0, 0

// 	for right := range s {
// 		cnt[s[right]-'A']++
// 		maxCnt = max(maxCnt, cnt[s[right]-'A'])

// 		for right-left+1-maxCnt > k {
// 			cnt[s[left]-'A']--
// 			left++
// 		}

// 	}
// 	return len(s) - left
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

func characterReplacement(s string, k int) (ans int) {
	// 窗口内字符个数
	cnts := [26]int{}
	// 窗口
	l, r, maxCnt, n := 0, 0, 1, len(s)
	ans = -1 << 31
	for r < n {
		cnts[s[r]-'A']++
		if cnts[s[r]-'A'] > maxCnt {
			maxCnt = cnts[s[r]-'A']
		}

		for l < r && r-l+1-maxCnt > k {
			cnts[s[l]-'A']--
			l++
			break
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
		r++
	}

	return
}
