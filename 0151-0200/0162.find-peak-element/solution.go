package findpeakelement

func findPeakElement(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
