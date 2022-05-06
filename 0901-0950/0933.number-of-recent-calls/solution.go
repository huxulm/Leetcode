package numberofrecentcalls

import "sort"

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

// 单调队列
// 时间: O(n) 空间: O(1)
func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	for len(this.q) > 0 && t-this.q[0] > 3000 {
		this.q = this.q[1:]
	}
	return len(this.q)
}

// 二分查找
// 时间: O(log(n)) 空间: O(1)
func (this *RecentCounter) Ping_1(t int) int {
	this.q = append(this.q, t)
	// 左边界 >= t - 3000
	return len(this.q) - sort.SearchInts(this.q, t-3000)
}
