package sqrtx

func mySqrt(x int) int {

	if x <= 1 {
		return x
	}

	// 二分法迭代
	l, r := 1, x
	ans := -1
	for l <= r {
		m := l + (r-l)>>1
		if m*m <= x {
			l = m + 1
			ans = m
		} else {
			r = m - 1
		}
	}
	return ans
}
