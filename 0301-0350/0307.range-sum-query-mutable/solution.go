package rangesumquerymutable

// 分块
// type NumArray struct {
// 	nums, sum []int
// 	size      int
// }

// func Constructor(nums []int) NumArray {
// 	n := len(nums)
// 	size := sort.Search(n, func(i int) bool { return (i+1)*(i+1) > n })
// 	// 向上取整
// 	size = (n-1)/size + 1
// 	sum := make([]int, size)
// 	for i := range nums {
// 		sum[i/size] += nums[i]
// 	}
// 	return NumArray{nums: nums, size: size, sum: sum}
// }

// func (na *NumArray) Update(index int, val int) {
// 	na.sum[index/na.size] += (val - na.nums[index])
// 	na.nums[index] = val
// }

// func (na *NumArray) SumRange(l int, r int) (ans int) {
// 	p, q := l/na.size, r/na.size
// 	// 同一分段
// 	if p == q {
// 		for i := l; i <= r; i++ {
// 			ans += na.nums[i]
// 		}
// 		return
// 	} else {
// 		for i := l; i < (p+1)*na.size; i++ {
// 			ans += na.nums[i]
// 		}
// 		// [p+1, q-1] 段
// 		for i := p + 1; i <= q-1; i++ {
// 			ans += na.sum[i]
// 		}
// 		for i := r; i >= q*na.size; i-- {
// 			ans += na.nums[i]
// 		}
// 		return
// 	}
// }

type NumArray struct {
	nums []int
	tree []int
}

// 树状数组
func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n+1)
	na := NumArray{tree: tree, nums: nums}
	for i := range nums {
		na.add(i+1, nums[i])
	}
	return na
}

func (na *NumArray) add(index int, val int) {
	for index < len(na.tree) {
		na.tree[index] += val
		index += index & -index
	}
}

func (na *NumArray) sum(index int) (sum int) {
	for index > 0 {
		sum += na.tree[index]
		index &= index - 1
	}
	return
}

func (na *NumArray) Update(index int, val int) {
	na.add(index+1, val-na.nums[index])
	na.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum(right+1) - this.sum(left)
}
