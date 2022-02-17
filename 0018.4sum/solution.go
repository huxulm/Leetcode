package sum

import "sort"

func nSumTarget(nums []int, n, start, target int) (ans [][]int) {
	sz := len(nums)
	// 至少是2sum且数组大小不小于n
	if n < 2 || sz < n {
		return
	}
	// base case: 2sum
	if n == 2 {
		lo, hi := start, sz-1
		for lo < hi {
			temp := nums[lo] + nums[hi]
			if temp == target {
				ans = append(ans, []int{nums[lo], nums[hi]})
				for lo++; lo < hi && nums[lo] == nums[lo-1]; lo++ {
				}
				for hi--; lo < hi && nums[hi] == nums[hi+1]; hi-- {
				}
			} else if temp > target {
				hi--
			} else {
				lo++
			}
		}
	} else if n > 2 {
		// 枚举第1个数
		for i := start; i < sz; i++ {
			// 跳过已经枚举过的
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			for _, subAns := range nSumTarget(nums, n-1, i+1, target-nums[i]) {
				ans = append(ans, append([]int{nums[i]}, subAns...))
			}
		}
	}
	return ans
}

func fourSum(nums []int, target int) [][]int {
	// 升序
	sort.Ints(nums)
	return nSumTarget(nums, 4, 0, target)
}
