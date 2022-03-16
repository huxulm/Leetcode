package pascalstriangle

func generate(numRows int) [][]int {
	tri := make([][]int, numRows)

	for i := range tri {
		tri[i] = make([]int, i+1)
		if i == 0 {
			tri[i][0] = 1
			continue
		}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tri[i][j] = tri[i-1][0]
				continue
			}
			tri[i][j] = tri[i-1][j-1] + tri[i-1][j]
		}
	}
	return tri
}
