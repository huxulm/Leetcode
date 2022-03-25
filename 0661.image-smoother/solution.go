package imagesmoother

func imageSmoother(img [][]int) (ans [][]int) {
	m, n := len(img), len(img[0])
	ans = make([][]int, m)
	for i := range img {
		ans[i] = make([]int, n)
		for j := range img[i] {
			sum := 0
			cnt := 0
			for r := -1; r <= 1; r++ {
				for c := -1; c <= 1; c++ {
					if x, y := r+i, c+j; x >= 0 && x < m && y >= 0 && y < n {
						sum += img[x][y]
						cnt++
					}
				}
			}
			ans[i][j] = sum / cnt
		}
	}

	return
}
