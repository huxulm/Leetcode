package sortarraybyparity

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	p1, p2 := 0, n-1

	// for p1 < p2 {
	// 	for p1 < p2 && nums[p1]&1 == 0 {
	// 		p1++
	// 	}
	// 	for p1 < p2 && nums[p2]&1 == 1 {
	// 		p2--
	// 	}
	// 	nums[p1], nums[p2] = nums[p2], nums[p1]
	// }

	for p1 < p2 {
		if nums[p1]&1 == 0 {
			p1++
			continue
		}
		if nums[p2]&1 == 1 {
			p2--
			continue
		}
		nums[p1], nums[p2] = nums[p2], nums[p1]
	}
	return nums
}
