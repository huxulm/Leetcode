package minimumfallingpathsum

func minFallingPathSum(A [][]int) (ans int) {
	m, n := len(A), len(A[0])
	var temp int
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			temp = A[i-1][j]
			for k := j - 1; k <= j+1; k++ {
				if k >= 0 && k < n {
					temp = min(temp, A[i-1][k])
				}
			}
			A[i][j] = temp + A[i][j]
		}
	}
	ans = A[m-1][0]
	for i := range A {
		ans = min(ans, A[m-1][i])
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
