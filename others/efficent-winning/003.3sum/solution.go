package sum

import (
	"sort"
)

// 分析:
// 要求找到 a + b + c = 0 的不重复三元组，可以先固定一位数，
// 然后求两数之和。要求不重复的三元组，需要考虑重复问题。
//
func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}

	// 升序
	sort.Ints(nums)

	for first := 0; first < n; first++ {
		// 需要和上次枚举的不同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		// 左右指针寻找两数之和 b+c = -a
		target := -nums[first]

		// c 对应的指针初始指向最右端
		third := n - 1

		// 枚举b
		for second := first + 1; second < third; second++ {
			// 需要和上次枚举的不同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 保证b指针在c指针左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			// 指针重合
			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ans
}
