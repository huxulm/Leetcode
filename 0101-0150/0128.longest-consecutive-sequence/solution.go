package longestconsecutivesequence

func longestConsecutive(nums []int) (ans int) {
	// start => len
	m := map[int]bool{}
	ans = 0
	for _, x := range nums {
		m[x] = true
	}
	for k := range m {
		if _, ok := m[k-1]; !ok {
			curLen := 1
			cur := k + 1
			ans = max(ans, curLen)
			for {
				if _, ok := m[cur]; ok {
					curLen++
					ans = max(curLen, ans)
					cur++
				} else {
					break
				}
			}
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
