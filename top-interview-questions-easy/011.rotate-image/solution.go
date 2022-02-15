package rotateimage

func rotate(matrix [][]int) {
	n := len(matrix)
	matrix_new := make([][]int, n)

	for i := range matrix_new {
		matrix_new[i] = make([]int, n)
	}

	for i := range matrix {
		for j, v := range matrix[i] {
			matrix_new[j][n-1-i] = v
		}
	}
	// 拷贝 matrix_new 矩阵每行的引用
	copy(matrix, matrix_new)
}
