package fibonaccinumber

func fib(n int) int {
	if n < 2 {
		return n
	}

	var f1, f2 int = 0, 1

	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}
