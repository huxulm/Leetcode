package maximumnumberofoccurrencesofasubstring

// 方法一: 枚举
func maxFreq(s string, maxLetters int, minSize int, maxSize int) (ans int) {

	m := map[string]int{} // 子串长度

	n := len(s)

	// 枚举子串左端点
	for i := range s {

		// 最短的子串 字符数不符合要求
		cnts := [26]bool{}
		tot := 0
		for j := i; j <= i+maxSize-1 && j < n; j++ {
			ch := s[j] - 'a'
			if !cnts[ch] {
				cnts[ch] = true
				tot++
			}
			if tot > maxLetters {
				break
			} else if j >= i+minSize-1 {
				m[s[i:j+1]]++
			}
		}

	}

	for _, v := range m {
		if v > ans {
			ans = v
		}
	}

	return
}

// 方法二: 可行性优化
func maxFreq1(s string, maxLetters int, minSize int, maxSize int) (ans int) {
	n := len(s)
	m := map[string]int{}

	// 在长度 [minSize, maxSize] 中所有满足条件的子串次数最多的一定不会超过
	// 长度 = minSize的子串个数
	for i := 0; i < n-minSize+1; i++ {

		// 枚举每个 左端点为 i 长度为 minSize 的子串
		cur := s[i : i+minSize]
		cnt := [26]bool{}
		tot := 0
		for _, ch := range cur {
			if !cnt[ch-'a'] {
				tot++
				cnt[ch-'a'] = true
			}
		}

		// 子串不同字符数判断
		if tot <= maxLetters {
			m[cur]++
			if m[cur] > ans {
				ans = m[cur]
			}
		}
	}
	return
}
