package permutationsequence

import "strconv"

func getPermutation(n int, k int) string {
	var (
		count int = 0
		nums      = make([]byte, n)
		track     = make([]byte, n)
		vis       = make([]bool, n)
		ans   string
	)

	// 初始化
	for i := range nums {
		nums[i] = '1' + byte(i)
	}
	backtrack(nums, &count, vis, track, &ans, k, 0)
	return ans
}

func backtrack(nums []byte, count *int, vis []bool, track []byte, ans *string, k, idx int) {

	if *count == k {
		return
	}

	if idx == len(vis) {
		*count++
		if *count == k {
			*ans = string(track)
		}
		return
	}

	for i := range nums {
		if !vis[i] {
			vis[i] = true
			track[idx] = nums[i]
			backtrack(nums, count, vis, track, ans, k, idx+1)
			vis[i] = false
		}
	}
}

func getPermutation1(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--

	ans := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k/factorial[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n-i]
	}
	return ans
}
