package rangesumquerymutable

import (
	"sort"
)

// 分块
type NumArray struct {
	nums, sum []int
	size      int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	size := sort.Search(n, func(i int) bool { return (i+1)*(i+1) > n })
	// 向上取整
	size = (n-1)/size + 1
	sum := make([]int, size)
	for i := range nums {
		sum[i/size] += nums[i]
	}
	return NumArray{nums: nums, size: size, sum: sum}
}

func (na *NumArray) Update(index int, val int) {
	na.sum[index/na.size] += (val - na.nums[index])
	na.nums[index] = val
}

func (na *NumArray) SumRange(l int, r int) (ans int) {
	p, q := l/na.size, r/na.size
	// 同一分段
	if p == q {
		for i := l; i <= r; i++ {
			ans += na.nums[i]
		}
		return
	} else {
		for i := l; i < (p+1)*na.size; i++ {
			ans += na.nums[i]
		}
		// [p+1, q-1] 段
		for i := p + 1; i <= q-1; i++ {
			ans += na.sum[i]
		}
		for i := r; i >= q*na.size; i-- {
			ans += na.nums[i]
		}
		return
	}
}
