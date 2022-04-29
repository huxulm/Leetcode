package gridgame

import "math"

func gridGame(g [][]int) (ans int64) {
	ans = math.MaxInt64

	left0 := 0
	for _, v := range g[0] {
		left0 += v
	}

	left1 := 0
	for i, v := range g[1] {
		left0 -= g[0][i]
		ans = min(ans, max(int64(left0), int64(left1)))
		left1 += v
	}
	return
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
