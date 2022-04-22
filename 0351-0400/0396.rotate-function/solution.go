package rotatefunction

func maxRotateFunction(nums []int) (ans int) {
	n := len(nums)

	sum := 0

	// F(K+1) = F(K) - A[n-1]*(n-1) + sum - A[n-1]
	F := 0

	for i := range nums {
		F += nums[i] * i
		sum += nums[i]
	}

	ans = F
	for i := 1; i < n; i++ {
		F = F - nums[(n-i)%n]*(n-1) + sum - nums[(n-i)%n]
		ans = max(ans, F)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
