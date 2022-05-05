// https://leetcode-cn.com/problems/subarray-product-less-than-k/
package subarrayproductlessthank

import "math"

// 方法一: 二分查找
// 思路:
// 对于固定i，二分查找出最大的j满足 [i,j] 内乘积小于k
// 考虑到乘积很容易溢出，所以对 [i,j] 内每一项取对数再求和来代替乘法运算
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	n := len(nums)
	prefix := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + math.Log(float64(nums[i-1]))
	}

	logK := math.Log(float64(k))

	for i := 0; i < n+1; i++ {
		lo, hi := i+1, n+1
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			// 不加 1e-9 的误差, 不能完全通过
			// 题意中 nums[i] <= 1000, log(1000)-log(999) < 0.0005
			if prefix[mid]-prefix[i]+1e-9 >= logK {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		ans += lo - i - 1
	}

	return
}

// 方法二: 双指针/滑动窗口
// func numSubarrayProductLessThanK1(nums []int, k int) (ans int) {
// 	n := len(nums)
// 	l, r := 0, 0
// 	mul := 1
// 	for r < n {
// 		mul *= nums[r]
// 		for mul >= k && l <= r {
// 			mul /= nums[l]
// 			l++
// 		}
// 		ans += r - l + 1
// 		r++
// 	}
// 	return ans
// }

func numSubarrayProductLessThanK1(nums []int, k int) (ans int) {
	n := len(nums)
	l, r, mul := 0, 0, 1
	for r < n {
		mul *= nums[r]
		for l <= r && mul >= k {
			mul /= nums[l]
			l++
		}
		ans += r - l + 1
		r++
	}
	return ans
}
