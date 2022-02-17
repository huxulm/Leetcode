package rotate

func rotate(nums []int, k int) {

	k %= len(nums)

	var reverse func(arr []int)
	reverse = func(arr []int) {
		n := len(arr)
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
	}

	reverse(nums)     //先反转全部的元素
	reverse(nums[:k]) // 反转前k个元素
	reverse(nums[k:]) // 反转剩余元素
}
