package plusone

func plusOne(digits []int) []int {
	var high = 0
	var n = len(digits)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			digits[i] += 1
		}
		v := digits[i] / 10
		digits[i] %= 10
		if v == 0 { // 不需要进位
			break
		} else {
			if i-1 >= 0 {
				digits[i-1] += v
			} else {
				high = v
			}
		}
	}

	if high == 0 {
		return digits
	} else {
		return append([]int{high}, digits...)
	}
}
