package adddigits

func addDigits(x int) (ans int) {
	if x < 10 {
		return x
	}
	for x > 0 {
		ans += x % 10
		x /= 10
	}

	return addDigits(ans)
}

func addDigits1(num int) int {
	for num >= 10 {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
	}
	return num
}

func addDigits2(num int) int {
	return (num-1)%9 + 1
}
