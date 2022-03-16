package addbinary

func addBinary(a string, b string) string {
	var arr_a, arr_b []byte
	n, m := len(a), len(b)
	if n > m {
		arr_a, arr_b = []byte(a), []byte(b)
	} else {
		arr_a, arr_b = []byte(b), []byte(a)
		n, m = m, n
	}

	// 最高位
	var high byte

	for i, j := n-1, m-1; i >= 0; i, j = i-1, j-1 {
		if j >= 0 {
			arr_a[i] = arr_a[i] + arr_b[j] - '0'
		}

		div := (arr_a[i] - '0') / 2
		arr_a[i] = '0' + (arr_a[i]-'0')%2
		// 进位
		if i >= 1 {
			arr_a[i-1] += div
		} else {
			high = div
		}
	}

	if high > 0 {
		return string(append([]byte{'0' + high}, arr_a...))
	}

	return string(arr_a)
}
