package minimumoperationstoconvertnumber

func minimumOperations(nums []int, start int, goal int) int {
	// i => goal 步数
	// memo[i] = 0 未搜索过
	memo := make([]int, 1001)

	q := []int{start}

	for step := 0; len(q) > 0; step++ {

		tmp := q
		q = nil

		for _, x := range tmp {

			if x == goal {
				return step
			}

			for _, v := range nums {
				a, b, c := x+v, x-v, x^v
				if a == goal || b == goal || c == goal {
					return step + 1
				}
				if a >= 0 && a <= 1000 && memo[a] == 0 {
					memo[a] = step + 1
					q = append(q, a)
				}
				if b >= 0 && b <= 1000 && memo[b] == 0 {
					memo[b] = step + 1
					q = append(q, b)
				}
				if c >= 0 && c <= 1000 && memo[c] == 0 {
					memo[c] = step + 1
					q = append(q, c)
				}
			}

		}
	}

	return -1
}
