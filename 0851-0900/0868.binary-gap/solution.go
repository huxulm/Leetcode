package binarygap

func binaryGap(n int) (ans int) {
	last := -1

	for i := 0; n > 0; i, n = i+1, n>>1 {
		if n&1 == 1 {
			if last != -1 {
				ans = max(i-last, ans)
			}
			last = i
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
