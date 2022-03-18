package maximumsumcircularsubarray

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 方法 1：邻接数组
// https://leetcode-cn.com/problems/maximum-sum-circular-subarray/solution/huan-xing-zi-shu-zu-de-zui-da-he-by-leetcode/
func maxSubarraySumCircular(A []int) int {
	N := len(A)

	ans, cur := A[0], A[0]
	for i := 1; i < N; i++ {
		cur = A[i] + max(cur, 0)
		ans = max(ans, cur)
	}

	// ans is the answer for 1-interval subarrays.
	// Now, let's consider all 2-interval subarrays.
	// For each i, we want to know
	// the maximum of sum(A[j:]) with j >= i+2

	// rightsums[i] = A[i] + A[i+1] + ... + A[N-1]
	rightsums := make([]int, N)
	rightsums[N-1] = A[N-1]
	for i := N - 2; i >= 0; i-- {
		rightsums[i] = rightsums[i+1] + A[i]
	}

	// maxright[i] = max_{j >= i} rightsums[j]
	maxright := make([]int, N)
	maxright[N-1] = A[N-1]
	for i := N - 2; i >= 0; i-- {

		maxright[i] = max(maxright[i+1], rightsums[i])
	}

	leftsum := 0
	for i := 0; i < N-2; i++ {
		leftsum += A[i]
		ans = max(ans, leftsum+maxright[i+2])
	}

	return ans
}
