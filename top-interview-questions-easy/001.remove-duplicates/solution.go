package removeduplicates

func removeDuplicates(nums []int) int {
	n := len(nums)

	if n <= 1 {
		return n
	}

	slow, fast := 0, 1

	for fast < n {
		for nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}
