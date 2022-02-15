package singlenumber

func singleNumber(nums []int) int {
	single := nums[0]
	for _, x := range nums[1:] {
		single ^= x
	}
	return single
}
