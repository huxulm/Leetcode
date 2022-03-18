package nthtribonaccinumber

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}

	var t0, t1, t2 int = 0, 1, 1

	for i := 3; i <= n; i++ {
		t0, t1, t2 = t1, t2, t0+t1+t2
	}

	return t2
}
