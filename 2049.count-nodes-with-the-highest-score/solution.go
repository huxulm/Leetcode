package countnodeswiththehighestscore

func countHighestScoreNodes(parents []int) (ans int) {
	n := len(parents)
	g := make([][]int, n)
	for i, p := range parents {
		if p == -1 {
			continue
		}
		g[p] = append(g[p], i)
	}

	maxScore, cnt := 1, 0

	var dfs func(node int) int

	dfs = func(node int) int {
		score, size := 1, n-1

		for _, ch := range g[node] {
			sz := dfs(ch)
			score *= sz
			size -= sz
		}

		if node != 0 {
			score *= size
		}

		if score == maxScore {
			cnt++
		} else if score > maxScore {
			maxScore = score
			cnt = 1
		}

		return n - size
	}
	dfs(0)
	return
}
