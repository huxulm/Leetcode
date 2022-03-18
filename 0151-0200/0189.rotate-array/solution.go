package rotatearray

func rotate(nums []int, k int) {
	var reverse func(arr []int)
	reverse = func(arr []int) {
		n := len(arr)
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
