package minimumsizesubarraysum

// 暴力: O(n^2)
func minSubArrayLen(target int, nums []int) (ans int) {
	ans = 1<<31 - 1
	n := len(nums)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= target && j-i+1 < ans {
				ans = j - i + 1
				break
			}
		}
	}
	if ans == 1<<31-1 {
		ans = 0
	}
	return
}

// 前缀和+二分优化
// 思路:
// 对于每一个数组起始位置i，找到最小的下标j使得：[i, j]区间的和大于等于target，
// 则用（j-i+1）更新答案
// 用sum数组表示nums的前缀和, sum[i] = nums[0] + ... + nums[i-1]
func minSubArrayLen1(target int, nums []int) (ans int) {
	ans = 1<<31 - 1
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for i := 0; i < n; i++ {
		// j 刚好是[i,j-1]的长度
		j := lowerBound(sum[i:], sum[i]+target)
		if j >= 0 && j-i < ans {
			ans = j - i
		}
	}
	if ans == 1<<31-1 {
		ans = 0
	}
	return
}

func lowerBound(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n
	for l < r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == 0 && nums[l] != target || l == n {
		return -1
	}
	return l
}
