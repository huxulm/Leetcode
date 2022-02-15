package movezeros

func moveZeroes(nums []int) {
	n := len(nums)
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
