package nrepeatedelementinsize2narray

func repeatedNTimes(nums []int) (ans int) {
	m := map[int]int{}

	for _, v := range nums {
		m[v]++
		if m[v] > 1 {
			ans = v
			break
		}
	}
	return
}

// 数学优化
func repeatedNTimes1(nums []int) int {
	for gap := 1; gap <= 3; gap++ {
		for i, num := range nums[:len(nums)-gap] {
			if num == nums[i+gap] {
				return num
			}
		}
	}
	return -1 // 不可能的情况
}
