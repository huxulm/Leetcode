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
