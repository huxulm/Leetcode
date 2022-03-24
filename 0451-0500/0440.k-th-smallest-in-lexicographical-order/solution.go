package kthsmallestinlexicographicalorder

func findKthNumber(n int, k int) int {
	cur := 1
	k--

	for k > 0 {
		steps := getSteps(cur, n)

		if steps <= k {
			k -= steps
			cur++
		} else {
			cur *= 10
			k--
		}
	}

	return cur
}

func getSteps(cur, n int) (steps int) {
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
