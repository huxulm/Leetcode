package findmissingobservations

func missingRolls(rolls []int, mean int, n int) (ans []int) {
	m := len(rolls)
	mSum := 0
	for i := range rolls {
		mSum += rolls[i]
	}

	nSum := mean*(m+n) - mSum

	if nSum > n*6 || nSum < n {
		return
	}

	// ans[i] 在[1,6]
	ans = make([]int, n)

	ava, remain := nSum/n, nSum%n

	for i := range ans {
		if remain > 0 {
			ans[i] = ava + 1
			remain--
		} else {
			ans[i] = ava
		}
	}

	return
}
