package squaresofasortedarray

func sortedSquares(nums []int) []int {
	n := len(nums)
	pivot := 0
	for i := range nums {
		if nums[i] < 0 {
			pivot = i
		}
		// 计算平方
		nums[i] *= nums[i]
	}

	p1, p2 := pivot, pivot+1
	ans := make([]int, n)

	index := 0
	for p1 >= 0 && p2 < n {

		if nums[p1] < nums[p2] {
			ans[index] = nums[p1]
			p1--
		} else {
			ans[index] = nums[p2]
			p2++
		}
		index++
	}
	for p1 >= 0 && p2 == n {
		ans[index] = nums[p1]
		index++
		p1--
	}
	for p2 < n && p1 == -1 {
		ans[index] = nums[p2]
		index++
		p2++
	}
	return ans
}
