package checkifnumberhasequaldigitcountanddigitvalue

func digitCount(num string) bool {
	n := len(num)
	needCnt := make([]int, n)
	cnt := make([]int, 10)
	for i := range num {
		v := int(num[i] - '0')
		needCnt[i] = v
		cnt[v]++
	}

	for i := range needCnt {
		if needCnt[i] != cnt[i] {
			return false
		}
	}

	return true
}
