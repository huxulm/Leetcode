package maximumsubarray

// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
func maxSubArray(nums []int) int {
	// nums[i]结尾的连续子数组最大和
	n := len(nums)
	ans := -1 << 31
	var dp int

	for i := 0; i < n; i++ {
		if i == 0 {
			dp = nums[i]
		} else {
			// dp = max(dp+nums[i], nums[i])
			if dp > 0 {
				dp = dp + nums[i]
			} else {
				dp = nums[i]
			}
		}
		// 更新答案
		if dp > ans {
			ans = dp
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
