package countnodeswiththehighestscore

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	// 构建树(邻接表表示)
	g := make([][]int, n)
	for i, p := range parents[1:] {
		g[p] = append(g[p], i+1)
	}

	// 最高分数，最高分数数量
	maxScore, cnt := 0, 0

	var dfs func(node int) int

	dfs = func(node int) int {

		score := 1
		size := 1

		for _, ch := range g[node] {
			sz := dfs(ch)
			score *= sz
			size += sz
		}

		if node != 0 {
			score *= n - size
		}

		if score == maxScore {
			cnt++
		} else if score > maxScore {
			maxScore = score
			cnt = 1
		}

		return size
	}
	dfs(0)
	return cnt
}

func countHighestScoreNodes1(parents []int) int {
	n := len(parents)
	matrix := make([][]int, 4)
	for i := 0; i < 4; i++ {
		matrix[i] = make([]int, n)
	}
	for _, p := range parents {
		if p >= 0 {
			matrix[0][p]++
		}
	}

	var cur []int
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			cur = append(cur, i)
			//matrix[1][i] = 1
		}
	}

	for len(cur) > 0 {
		var next []int
		for _, c := range cur {
			p := parents[c]
			if p == -1 {
				continue
			}
			matrix[0][p]--
			matrix[1][p] += matrix[1][c] + 1
			if matrix[2][p] > 0 {
				matrix[3][p] = matrix[1][c] + 1
			} else {
				matrix[2][p] = matrix[1][c] + 1
			}
			if matrix[0][p] == 0 {
				next = append(next, p)
			}
		}
		cur = next
	}

	max := n - 1
	for i := 0; i < n; i++ {
		tmp := matrix[2][i]
		if tmp == 0 {
			tmp = 1
		}
		if matrix[3][i] > 1 {
			tmp *= matrix[3][i]
		}
		p := n - matrix[3][i] - matrix[2][i] - 1
		if p > 1 {
			tmp *= p
		}
		matrix[0][i] = tmp
		if max < tmp {
			max = tmp
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if matrix[0][i] == max {
			res++
		}
	}

	return res
}
