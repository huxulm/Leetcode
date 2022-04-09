package minimumswapstogroupall1stogetherii

func minSwaps(nums []int) (ans int) {
	n := len(nums)
	// 1的个数
	cnt1 := 0

	// 区间内0的个数
	pre := make([]int, n+1)
	for i := range nums {
		pre[i+1] = pre[i]
		if nums[i] == 0 {
			pre[i+1] += 1
		} else {
			cnt1++
		}
	}

	if cnt1 <= 1 || cnt1 == n {
		return
	}

	ans = 1<<31 - 1

	for i := 0; i < n; i++ {
		end := (i + cnt1 - 1) % n
		if end > i {
			ans = min(ans, pre[end+1]-pre[i])
		} else if end < i {
			ans = min(ans, pre[end+1]+pre[n]-pre[i])
		}
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
