package findthewinnerofthecirculargame

func findTheWinner(n int, k int) int {
	if k == 1 {
		return n
	}
	q := make([]int, n)
	for i := range q {
		q[i] = i + 1
	}
	idx := 0
	for len(q) > 1 {
		tmp := q
		q = nil
		for _, v := range tmp {
			idx++
			if idx%k != 0 {
				q = append(q, v)
			}
		}
	}
	return q[0]
}
