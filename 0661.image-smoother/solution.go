package imagesmoother

// 方法一: 朴素解法
// 时间: O(m*n)
// 空间: O(m*n)
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

// 方法二: 二维前缀和
func imageSmoother1(img [][]int) (ans [][]int) {
	m, n := len(img), len(img[0])
	preSum := make([][]int, m+10)
	preSum[0] = make([]int, n+10)

	for i := 1; i <= m; i++ {
		preSum[i] = make([]int, n+10)
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + img[i-1][j-1]
		}
	}

	ans = make([][]int, m)

	for i := range img {
		ans[i] = make([]int, n)
		for j := range img[i] {
			a, b := max(0, i-1), max(0, j-1)
			c, d := min(i+1, m-1), min(j+1, n-1)
			cnts := (c - a + 1) * (d - b + 1)
			tot := preSum[c+1][d+1] - preSum[c+1][b] - preSum[a][d+1] + preSum[a][b]
			ans[i][j] = tot / cnts
		}
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
