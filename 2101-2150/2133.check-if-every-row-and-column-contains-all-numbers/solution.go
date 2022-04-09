package checkifeveryrowandcolumncontainsallnumbers

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	// 行
	for i := range matrix {
		m := map[int]bool{}
		for _, v := range matrix[i] {
			if !m[v] && v > 0 && v <= n {
				m[v] = true
			} else {
				return false
			}
		}
	}

	// 列
	for j := 0; j < n; j++ {
		m := map[int]bool{}
		for i := 0; i < n; i++ {
			if v := matrix[i][j]; !m[v] && v > 0 && v <= n {
				m[v] = true
			} else {
				return false
			}
		}
	}

	return true
}
