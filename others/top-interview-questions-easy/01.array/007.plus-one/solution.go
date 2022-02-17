package plusone

func plusOne(digits []int) []int {
	n := len(digits)

	// 标志进位
	flag := 0
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			digits[i]++
		}
		digits[i] += flag
		flag = digits[i] / 10
		digits[i] %= 10
		if flag == 0 {
			break
		}
	}
	if flag == 0 {
		return digits
	} else {
		return append([]int{flag}, digits...)
	}
}
