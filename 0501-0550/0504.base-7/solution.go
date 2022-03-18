package base7

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var sign bool
	if num < 0 {
		sign = true
		num = -num
	}

	ans := []byte{}
	for num > 0 {
		ans = append(ans, byte(num%7+'0'))
		num /= 7
	}

	if sign {
		ans = append(ans, '-')
	}

	for i, j := 0, len(ans); i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return string(ans)
}
