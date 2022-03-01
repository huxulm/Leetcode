package zigzagconversion

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}
	res := make([]byte, n)
	index := 0
	for i := 0; i < numRows; i++ {

		for j, k := i, 0; j < n; {
			res[index] = s[j]
			index++
			if i != 0 && i != numRows-1 {
				j2 := 2*(k+1)*(numRows-1) - i
				if j2 >= 0 && j2 < n {
					res[index] = s[j2]
					index++
				}
			}
			k++
			j = i + 2*k*(numRows-1)
		}
	}
	return string(res)
}

func convert1(s string, r int) string {
	n := len(s)
	if r == 1 {
		return s
	}

	ans := make([]byte, 0, n)

	t := 2*r - 2
	for i := 0; i < r; i++ {
		for j := 0; i+j < n; j = j + t {
			ans = append(ans, s[i+j])          // 当前周期的第一个字符
			if i > 0 && i < r-1 && j+t-i < n { // (i+j+t)-2*i
				ans = append(ans, s[j+t-i])
			}
		}
	}

	return string(ans)
}
