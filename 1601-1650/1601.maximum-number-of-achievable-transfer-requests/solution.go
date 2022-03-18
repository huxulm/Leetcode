package maximumnumberofachievabletransferrequests

func maximumRequests(n int, requests [][]int) (ans int) {
	m := len(requests)

	// 每栋楼的进出量
	var delta = make([]int, n)
	var cnt, zero = 0, n
	// 选择第[idx]个选择
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == m {
			if zero == n && cnt > ans {
				ans = cnt
			}
			return
		}

		// 不选择
		dfs(idx + 1)

		// 选择
		z := zero
		cnt++
		r := requests[idx]
		x, y := r[0], r[1]
		if delta[x] == 0 {
			zero--
		}
		delta[x]--
		if delta[x] == 0 {
			zero++
		}
		if delta[y] == 0 {
			zero--
		}
		delta[y]++
		if delta[y] == 0 {
			zero++
		}
		dfs(idx + 1)
		delta[y]--
		delta[x]++
		cnt--
		zero = z

	}

	dfs(0)
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
