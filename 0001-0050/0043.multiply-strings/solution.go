package multiplystrings

func multiply(num1 string, num2 string) string {
	// 结果恒为0
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1, n2 := len(num1), len(num2)
	N := n1 + n2

	var res = make([]byte, N)

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			res[i+j] += (num1[n1-1-i] - '0') * (num2[n2-1-j] - '0')
			res[i+j+1] += res[i+j] / 10
			res[i+j] = res[i+j] % 10
		}
	}

	if res[N-1] == 0 {
		res = res[0 : N-1]
	}

	var s string
	for i := len(res) - 1; i >= 0; i-- {
		s += string([]byte{'0' + res[i]})
	}
	return s
}
