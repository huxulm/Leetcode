package intersect

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	p1, p2 := 0, 0

	sort.Ints(nums1)
	sort.Ints(nums2)

	ans := []int{}

	for p1 < n1 && p2 < n2 {

		if nums1[p1] == nums2[p2] {
			ans = append(ans, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		}

	}

	return ans
}
