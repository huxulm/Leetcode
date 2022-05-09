package distringmatch

func diStringMatch(s string) (ans []int) {
	n := len(s)
	ans = make([]int, 0, n)
	min, max := 0, n
	for i := range s {
		if s[i] == 'I' {
			ans = append(ans, min)
			min++
		} else if s[i] == 'D' {
			ans = append(ans, max)
			max--
		}
	}
	ans = append(ans, min)
	return
}
