package mostbeautifulitemforeachquery

import "sort"

func maximumBeauty(items [][]int, queries []int) (ans []int) {
	// 按价格升序
	sort.Slice(items, func(i, j int) bool { return items[i][0] < items[j][0] })

	n := len(items)

	// 更新 items[i] 最大美丽值为 [i, j] 的最大美丽值
	for i := 1; i < n; i++ {
		if items[i-1][1] > items[i][1] {
			items[i][1] = items[i-1][1]
		}
	}

	ans = make([]int, len(queries))
	for i := range ans {
		x := queries[i]
		// [-1, n-1]
		j := sort.Search(n, func(mid int) bool { return items[mid][0] > x }) - 1
		if j == -1 {
			continue
		} else {
			ans[i] = items[j][1]
		}
	}
	return
}
