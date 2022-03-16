package rotateimage

func rotate(matrix [][]int) {
	// n * n
	// [1, 2, 3]     [7, 4, 1]
	// [4, 5, 6] ==> [8, 5, 2]
	// [7, 8, 9]     [9, 6, 3]

	n := len(matrix)

	// 依次转动每一层元素,由外向内
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			// top, right, bottom, left = left, top, right, bottom
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}
