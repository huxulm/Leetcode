package maximumwhitetilescoveredbyacarpet

import (
	"sort"
)

func maximumWhiteTiles(tiles [][]int, carpetLen int) (ans int) {
	// 左端点升序
	sort.Slice(tiles, func(i, j int) bool { return tiles[i][0] < tiles[j][0] })

	n := len(tiles)

	// 前缀和计算区间和
	var sum []int = make([]int, n+1)
	for i := range tiles {
		sum[i+1] = sum[i] + tiles[i][1] - tiles[i][0] + 1
	}

	for i := range tiles {
		maxPos := tiles[i][0] + carpetLen - 1
		// 找到最大的 <= maxPos的左端点
		// 最小的 >= maxPos - 1  <=> 最大的 <= maxPos
		j := sort.Search(n, func(i int) bool { return tiles[i][0] > maxPos }) - 1
		// n-1 >=j >= i
		ans = max(ans, sum[j]-sum[i]+min(maxPos, tiles[j][1])-tiles[j][0]+1)
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
