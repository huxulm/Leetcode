package maximumtotalimportanceofroads

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	g := make([][]int, n)

	for _, edge := range roads {
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	sort.Slice(g, func(i, j int) bool {
		return len(g[i]) > len(g[j])
	})

	var w int64

	for i := range g {
		if len(g[i]) == 0 {
			continue
		}
		w += int64(len(g[i]) * n)
		n -= 1
	}

	return w
}
