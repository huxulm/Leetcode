package removeduplicatesfromsortedarray

// 12[222567]
// 12[22567]|2
// 12[2567]|22
// 12[567]|222

// func removeDuplicates(nums []int) int {

// 	l := len(nums)
// 	dups := 0
// 	for i := 0; i < l-dups; i++ {
// 		for j := i + 1; j < l-dups; {
// 			// 如果nums[i] == nums[j], 将nums[j]移动到数组末尾
// 			if nums[i] == nums[j] {
// 				dups += 1
// 				for k := j; k <= l-dups-1; k++ {
// 					nums[k], nums[k+1] = nums[k+1], nums[k]
// 				}
// 			} else {
// 				j++
// 			}
// 		}
// 	}

// 	return len(nums) - dups
// }

// 有序数组去重: 快慢指针
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var slow, fast = 0, 0

	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			// 维护 nums[0..slow] 无重复
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
