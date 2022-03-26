package topkfrequentelements

import "container/heap"

type hp [][]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) (ans []int) {
	m := map[int]int{}

	for _, x := range nums {
		m[x]++
	}

	h := &hp{}
	for k, v := range m {
		heap.Push(h, []int{v, k})
	}

	ans = make([]int, k)

	for i := range ans {
		v := heap.Pop(h).([]int)
		ans[i] = v[1]
	}

	return
}
