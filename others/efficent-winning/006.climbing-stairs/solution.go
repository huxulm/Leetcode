package climbingstairs

func climbStairs(n int) int {
	if n < 2 {
		return n
	}

	var f1, f2 int = 2, 3

	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f1
}
