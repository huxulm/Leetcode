package simplifiedfractions

import "fmt"

func simplifiedFractions(n int) []string {
	ans := []string{}

	var isSimplified func(x, y int) bool

	isSimplified = func(x, y int) bool {
		for y != 0 {
			x, y = y, x%y
		}
		return x == 1
	}

	// 分母
	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			if isSimplified(i, j) {
				ans = append(ans, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}

	return ans
}
