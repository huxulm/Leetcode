package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	var (
		m, n                     = len(matrix), len(matrix[0])
		index                    = 0
		left, right, top, bottom = 0, n - 1, 0, m - 1
		ans                      = make([]int, m*n)
	)

	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			ans[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			ans[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for col := right - 1; col > left; col-- {
				ans[index] = matrix[bottom][col]
				index++
			}
			for row := bottom; row > top; row-- {
				ans[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}

	return ans
}
