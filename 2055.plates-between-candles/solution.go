package platesbetweencandles

import "runtime/debug"

func init() {
	debug.SetGCPercent(100)
}

func platesBetweenCandles(s string, queries [][]int) (ans []int) {
	n := len(s)
	// left[i] 表示 i 左侧第一个蜡烛位置
	left := make([]int, n)
	//  表示 s[:i] 中盘子的个数
	preSum := make([]int, n+1)
	for i, l := 0, -1; i < n; i++ {
		preSum[i+1] = preSum[i]
		if s[i] == '|' {
			l = i
		} else {
			preSum[i+1]++
		}
		left[i] = l
	}

	// right[i] 表示 i 右侧最近蜡烛位置
	right := make([]int, n)
	for i, r := n-1, -1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}

	ans = make([]int, len(queries))

	for i := range queries {
		l, r := right[queries[i][0]], left[queries[i][1]]
		if l >= 0 && r >= 0 && l < r {
			ans[i] = preSum[r+1] - preSum[l]
		}
	}

	return
}
