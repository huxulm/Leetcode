package luckynumbersinamatrix

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	minRow, maxCol := make([]int, m), make([]int, n)

	for i := 0; i < m; i++ {
		minRow[i] = matrix[i][0]
		for j := 0; j < n; j++ {
			minRow[i] = min(minRow[i], matrix[i][j])
			maxCol[j] = max(maxCol[j], matrix[i][j])
		}
	}

	ans := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == minRow[i] && matrix[i][j] == maxCol[j] {
				ans = append(ans, matrix[i][j])
			}
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
