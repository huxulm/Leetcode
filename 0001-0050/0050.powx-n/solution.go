package powxn

// 快速幂 + 迭代
// x => x^2 => x^4 => x^9 => x^19 => x^38 => x^77
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	ans := 1.0

	xContribute := x

	for n > 0 {
		if n%2 == 1 {
			ans *= xContribute
		}
		xContribute *= xContribute
		n >>= 1
	}
	return ans
}
