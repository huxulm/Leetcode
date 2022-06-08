package flipstringtomonotoneincreasing

func minFlipsMonoIncr(s string) (ans int) {
	n := len(s)
	preSumZero := make([]int, n+1)
	for i := range s {
		if s[i] == '0' {
			preSumZero[i+1] = preSumZero[i] + 1
		} else {
			preSumZero[i+1] = preSumZero[i]
		}
	}

	// 最多所有的 0 全部反转
	ans = preSumZero[n]
	// 枚举每个 i， [0,i] 变 0 [i+1..n-1] 变1
	for i := 0; i < n; i++ {
		// 左侧1个数 + 右侧0个数
		need := i + 1 - preSumZero[i+1] + preSumZero[n] - preSumZero[i+1]
		if need < ans {
			ans = need
		}
	}
	return
}
