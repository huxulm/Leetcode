package pascalstriangleii

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)

	// 开头
	row[0] = 1

	if rowIndex == 0 {
		return row
	}

	preRow := getRow(rowIndex - 1)

	for i := 0; i < rowIndex; i++ {
		row[i+1] = preRow[i] + preRow[i+1]
	}

	// 结尾
	row[rowIndex] = 1
	return row
}

// 迭代
func getRow1(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	var prev, cur []int = []int{1}, nil
	for i := 1; i <= rowIndex; i++ {
		n := i + 1
		cur = make([]int, n)
		for j := 0; j < n; j++ {
			if j == 0 || j == n-1 {
				cur[j] = 1
				continue
			}
			cur[j] = prev[j-1] + prev[j]
		}
		prev, cur = cur, nil
	}
	return prev
}
