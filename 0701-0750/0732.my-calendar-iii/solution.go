package mycalendariii

import "github.com/emirpasic/gods/trees/redblacktree"

// import "sort"

// type MyCalendarThree struct {
// 	l, r     int // 维护最多集合的交集区间
// 	maxBooks int
// 	Q        [][]int // 按区间左端点排序
// }

// func Constructor() MyCalendarThree {
// 	return MyCalendarThree{}
// }

// // <= 400 次
// func (this *MyCalendarThree) Book(start int, end int) int {
// 	// 返回最大预订次数
// 	// 返回最多交集的子集合 数量
// 	this.Q = append(this.Q, []int{start, end})

// 	if len(this.Q) == 1 {
// 		this.l, this.r = start, end
// 		this.maxBooks = 1
// 		return 1
// 	}

// 	// O(nlogn)
// 	sort.Slice(this.Q, func(i, j int) bool {
// 		return this.Q[i][0] < this.Q[j][0] || (this.Q[i][0] == this.Q[j][0] && this.Q[i][1] <= this.Q[j][1])
// 	})

// 	maxBooks := -1 << 31
// 	n := len(this.Q)
// 	for i := range this.Q {
// 		maxRight := this.Q[i][1]
// 		// 最大的小于等于 maxRight的左边界位置
// 		rIndex := sort.Search(n, func(m int) bool {
// 			return this.Q[m][0] > maxRight
// 		}) - 1
// 		if cnt := rIndex - i + 1; cnt > 0 && cnt > maxBooks {
// 			maxBooks = cnt
// 		}
// 	}
// 	return maxBooks
// }

type MyCalendarThree struct {
	*redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{redblacktree.NewWithIntComparator()}
}

func (t MyCalendarThree) add(x, delta int) {
	if val, ok := t.Get(x); ok {
		delta += val.(int)
	}
	t.Put(x, delta)
}

func (t MyCalendarThree) Book(start, end int) (ans int) {
	t.add(start, 1)
	t.add(end, -1)

	maxBook := 0
	for it := t.Iterator(); it.Next(); {
		maxBook += it.Value().(int)
		if maxBook > ans {
			ans = maxBook
		}
	}
	return
}
