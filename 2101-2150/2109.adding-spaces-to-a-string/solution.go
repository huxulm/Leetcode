package addingspacestoastring

func addSpaces(s string, spaces []int) string {
	n, m := len(s), len(spaces)

	ans := make([]byte, m+n)

	p1, p2 := 0, 0

	for _, end := range spaces {

		for idx := p2; idx < end; idx++ {
			ans[p1] = s[idx]
			p1++
			p2++
		}
		ans[p1] = ' '
		p1++
	}

	for p2 < n {
		ans[p1] = s[p2]
		p1++
		p2++
	}

	return string(ans)
}

func addSpaces1(s string, spaces []int) string {
	last := 0
	var res []byte
	for i := range spaces {
		res = append(res, s[last:spaces[i]]...)
		res = append(res, ' ')
		last = spaces[i]
	}
	res = append(res, s[last:]...)
	return string(res)
}
