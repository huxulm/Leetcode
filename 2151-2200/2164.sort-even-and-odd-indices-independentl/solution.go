package sortevenandoddindicesindependentl

import (
	"sort"
)

func sortEvenOdd(nums []int) []int {

	n := len(nums)

	// 奇/偶
	var even, odd []int

	odd = make([]int, n/2)

	if n%2 == 0 {
		even = make([]int, n/2)
	} else {
		even = make([]int, n/2+1)
	}

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			even[(i-1)/2] = nums[i-1]
		} else {
			odd[i/2-1] = nums[i-1]
		}
	}

	// 非递增
	sort.Slice(even, func(i, j int) bool {
		return even[i] <= even[j]
	})

	// 非递减
	sort.Slice(odd, func(i, j int) bool {
		return odd[i] >= odd[j]
	})

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			nums[i-1] = even[(i-1)/2]
		} else {
			nums[i-1] = odd[i/2-1]
		}
	}

	return nums
}
