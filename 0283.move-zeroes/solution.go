package movezeroes

func moveZeroes(nums []int) {
	slow, fast, n := 0, 0, len(nums)

	for fast < n {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
