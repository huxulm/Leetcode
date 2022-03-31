package findserversthathandledmostnumberofrequests

import (
	"container/heap"
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func busiestServers(k int, arrival, load []int) (ans []int) {
	available := redblacktree.NewWithIntComparator()
	for i := 0; i < k; i++ {
		available.Put(i, nil)
	}
	busy := hp{}
	requests := make([]int, k)
	maxRequest := 0
	for i, t := range arrival {
		for len(busy) > 0 && busy[0].end <= t {
			available.Put(busy[0].id, nil)
			heap.Pop(&busy)
		}
		if available.Size() == 0 {
			continue
		}
		node, _ := available.Ceiling(i % k)
		if node == nil {
			node = available.Left()
		}
		id := node.Key.(int)
		requests[id]++
		if requests[id] > maxRequest {
			maxRequest = requests[id]
			ans = []int{id}
		} else if requests[id] == maxRequest {
			ans = append(ans, id)
		}
		heap.Push(&busy, pair{t + load[i], id})
		available.Remove(id)
	}
	return
}

// type pair struct{ end, id int }
// type hp []pair
// func (h hp) Len() int            { return len(h) }
// func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
// func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
// func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 模拟+双优先队列
func busiestServers1(k int, arrival, load []int) (ans []int) {
	available := hi{make([]int, k)}
	for i := 0; i < k; i++ {
		available.IntSlice[i] = i
	}
	busy := hp{}
	requests := make([]int, k)
	maxRequest := 0
	for i, t := range arrival {
		for busy.Len() > 0 && busy[0].end <= t {
			heap.Push(&available, i+((busy[0].id-i)%k+k)%k) // 保证得到的是一个不小于 i 的且与 id 同余的数
			heap.Pop(&busy)
		}

		if available.Len() == 0 {
			continue
		}

		id := heap.Pop(&available).(int) % k
		requests[id]++

		if maxRequest < requests[id] {
			maxRequest = requests[id]
			ans = []int{id}
		} else if maxRequest == requests[id] {
			ans = append(ans, id)
		}

		heap.Push(&busy, pair{t + load[i], id})
	}
	return
}

type hi struct{ sort.IntSlice }

func (h *hi) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hi) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type pair struct{ end, id int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
