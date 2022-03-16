package medianoftwosortedarrays

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/comments/36497
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	left, right := (m+n+1)/2, (m+n+2)/2
	return float64(findKth(nums1, 0, nums2, 0, left)+findKth(nums1, 0, nums2, 0, right)) / 2.0
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1] //nums1为空数组
	}
	if j >= len(nums2) {
		return nums1[i+k-1] //nums2为空数组
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	var midVal1, midVal2 int
	if i+k/2-1 < len(nums1) {
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = 1<<31 - 1
	}
	if j+k/2-1 < len(nums2) {
		midVal2 = nums2[j+k/2-1]
	}
	if midVal1 < midVal2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKth(nums1, i, nums2, j+k/2, k-k/2)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
