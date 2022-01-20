package sum

import "sort"

func twoSumTarget(nums []int, start, target int) [][]int {
	sort.Ints(nums)
	lo, hi := start, len(nums)-1
	ans := make([][]int, 0)
	for lo < hi {
		left, right := nums[lo], nums[hi]
		sum := left + right
		if sum == target {
			ans = append(ans, []int{left, right})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	ans := make([][]int, 0)
	n := len(nums)

	// 穷举threeSum第一个数字
	for i := 0; i < n; i++ {
		tuples := twoSumTarget(nums, i+1, -nums[i])
		for _, tuple := range tuples {
			tuple = append([]int{nums[i]}, tuple...)
			ans = append(ans, tuple)
		}
		// 移除第一个数字重复的情况
		for ; i < n-1 && nums[i] == nums[i+1]; i++ {
		}
	}

	return ans
}

// LC O(n^2)
func threeSum1(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > 1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
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
