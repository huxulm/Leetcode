package countwordsobtainedafteraddingaletter

func wordCount(startWords []string, targetWords []string) (ans int) {
	var hash = func(s string) (h int) {
		for _, ch := range s {
			h |= 1 << (ch - 'a')
		}
		return
	}

	var m = make(map[int]struct{})

	for _, s := range startWords {
		v := hash(s)

		for i := 0; i < 26; i++ {
			if v>>i&1 == 0 {
				m[(1<<i)|v] = struct{}{}
			}
		}
	}

	for _, s := range targetWords {
		if _, ok := m[hash(s)]; ok {
			ans++
		}
	}

	return
}

// endlesscheng
// https://leetcode-cn.com/problems/count-words-obtained-after-adding-a-letter/solution/ni-xiang-si-wei-wei-yun-suan-ha-xi-biao-l4153/
func wordCount1(startWords, targetWords []string) (ans int) {
	has := map[int]bool{}
	for _, word := range startWords {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		has[mask] = true
	}
	for _, word := range targetWords {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		for _, ch := range word {
			// golang中的位清除: &^
			// v := mask
			// v &^= 1 << (ch - 'a')
			if has[mask^(1<<(ch-'a'))] { // 去掉这个字符
				ans++
				break
			}
		}
	}
	return
}
