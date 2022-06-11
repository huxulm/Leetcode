package reconstructa2rowbinarymatrix

func reconstructMatrix(upper int, lower int, colsum []int) (ans [][]int) {
	n := len(colsum)
	ans = make([][]int, 2)
	ans[0] = make([]int, n)
	ans[1] = make([]int, n)

	for i, v := range colsum {
		if v == 0 {
			continue
		} else {
			if v == 1 {
				if upper == 0 || lower == 0 {
					return [][]int{}
				}
				if upper >= lower {
					ans[0][i] = 1
					upper--
				} else {
					if lower == 0 {
						return [][]int{}
					}
					ans[1][i] = 1
					lower--
				}
			} else {
				if upper < 1 || lower < 1 {
					return [][]int{}
				}
				ans[0][i], ans[1][i] = 1, 1
				upper--
				lower--
			}
		}
	}
	if upper != 0 || lower == 0 {
		ans = [][]int{}
	}
	return
}
