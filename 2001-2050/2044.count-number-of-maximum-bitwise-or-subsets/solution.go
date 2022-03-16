package countnumberofmaximumbitwiseorsubsets

func countMaxOrSubsets(nums []int) (ans int) {
	n := len(nums)
	// max_or <= 65535
	max_or := 0
	for i := range nums {
		max_or |= nums[i]
	}

	for mask := 1; mask < 1<<n; mask++ {
		subsetsOr := 0
		for j := n - 1; j >= 0; j-- {
			if mask&(1<<j) > 0 {
				subsetsOr |= nums[j]
			}
			if subsetsOr == max_or {
				break
			}
		}
		if subsetsOr == max_or {
			ans++
		}
	}
	return
}
