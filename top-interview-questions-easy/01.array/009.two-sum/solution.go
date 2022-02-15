package twosum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := range nums {
		if _, ok := m[target-nums[i]]; ok {
			return []int{m[target-nums[i]], i}
		}
		m[nums[i]] = i
	}
	return nil
}
