package shortestdistancetoacharacter

func shortestToChar(s string, c byte) (ans []int) {
	n := len(s)
	ans = make([]int, n)

	last := -1
	for i := range ans {
		if s[i] == c {
			last = i
			ans[i] = 0
		} else {
			if last == -1 {
				ans[i] = 1<<31 - 1
			} else {
				ans[i] = i - last
			}
		}
	}

	last = -1
	for i := range ans {
		if s[n-i-1] == c {
			last = n - i - 1
			ans[n-i-1] = 0
		} else {
			if last != -1 {
				ans[n-i-1] = min(ans[n-i-1], last-n+i+1)
			}
		}
	}

	return
}

func shortestToChar1(s string, c byte) (ans []int) {
	n := len(s)
	ans = make([]int, n)

	last := -n
	for i := range ans {
		if s[i] == c {
			last = i
			ans[i] = 0
		}
		ans[i] = i - last
	}

	last = 2 * n
	for i := range ans {
		if s[n-i-1] == c {
			last = n - i - 1
		}
		ans[n-i-1] = min(ans[n-i-1], last-n+i+1)
	}

	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
