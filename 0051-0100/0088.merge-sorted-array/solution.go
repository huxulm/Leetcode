package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, x := m-1, n-1, m+n-1; x >= 0; x-- {
		if i >= 0 && j >= 0 {
			if nums1[i] < nums2[j] {
				nums1[x] = nums2[j]
				j--
			} else {
				nums1[x] = nums1[i]
				i--
			}
		} else if j < 0 && i >= 0 {
			nums1[x] = nums1[i]
			i--
		} else if i < 0 && j >= 0 {
			nums1[x] = nums2[j]
			j--
		}
	}
}
