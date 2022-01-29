package spiralmatrixii

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	var (
		idx                      = 0
		left, right, top, bottom = 0, n - 1, 0, n - 1
	)
	for left <= right && top <= bottom {
		// top
		for col := left; col <= right; col++ {
			ans[top][col] = idx + 1
			idx++
		}
		// right
		for row := top + 1; row <= bottom; row++ {
			ans[row][right] = idx + 1
			idx++
		}
		// bottom and left
		if left < right && top < bottom {
			// bottom
			for col := right - 1; col > left; col-- {
				ans[bottom][col] = idx + 1
				idx++
			}
			// left
			for row := bottom; row > top; row-- {
				ans[row][left] = idx + 1
				idx++
			}
		}
		left++
		right--
		top++
		bottom--
	}

	return ans
}
