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
