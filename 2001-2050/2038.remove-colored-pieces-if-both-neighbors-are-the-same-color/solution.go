package removecoloredpiecesifbothneighborsarethesamecolor

// 方法一: 模拟+暴力枚举，超时
// 分析: 复杂度 O(n^2)
func winnerOfGame(s string) (aliceWin bool) {
	// alice
	cur := 0
	var bobsWin bool
	for len(s) > 2 {
		for i := 1; i < len(s)-1; i++ {
			if cur == 0 {
				// alice
				if s[i] == 'A' && s[i-1] == s[i] && s[i] == s[i+1] {
					s = s[:i] + s[i+1:]
					aliceWin = true
					bobsWin = false
					break
				}
			} else {
				if s[i] == 'B' && s[i-1] == s[i] && s[i] == s[i+1] {
					s = s[:i] + s[i+1:]
					aliceWin = false
					bobsWin = true
					break
				}
			}
		}

		if (cur == 0 && !aliceWin) || (cur == 1 && !bobsWin) {
			return
		}

		cur = (cur + 1) % 2
	}

	return
}

func winnerOfGame1(s string) (aliceWin bool) {
	cnt1, cnt2 := 0, 0
	for i := 1; i < len(s)-1; i++ {
		if s[i] == 'A' && s[i-1] == s[i] && s[i] == s[i+1] {
			cnt1++
		}
		if s[i] == 'B' && s[i-1] == s[i] && s[i] == s[i+1] {
			cnt2++
		}
	}
	return cnt1 > cnt2
}
