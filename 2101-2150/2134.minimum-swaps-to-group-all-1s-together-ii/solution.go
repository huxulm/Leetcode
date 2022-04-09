package minimumswapstogroupall1stogetherii

// 方法一: 前缀和
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

// 方法二: 滑动窗口
func minSwaps1(nums []int) (ans int) {
	n := len(nums)
	cntOne := 0

	for i := range nums {
		if nums[i] == 1 {
			cntOne++
		}
	}

	if cntOne <= 1 || cntOne == n {
		return
	}

	ans = 1<<31 - 1
	cntZero := 0

	for i := 0; i < cntOne-1; i++ {
		if nums[i] == 0 {
			cntZero++
		}
	}

	for i := 0; i < n; i++ {
		end := (i + cntOne - 1) % n
		if nums[end] == 0 {
			cntZero++
		}

		ans = min(ans, cntZero)

		if nums[i] == 0 {
			cntZero--
		}
	}

	return
}
