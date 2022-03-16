package findminimuminrotatedsortedarray

func findMin(nums []int) (ans int) {
	n := len(nums)
	lo, hi := 0, n-1
	for lo < hi {
		pivot := lo + (hi-lo)>>1
		if nums[pivot] < nums[hi] {
			hi = pivot
		} else {
			lo = pivot + 1
		}
	}
	return nums[lo]
}
