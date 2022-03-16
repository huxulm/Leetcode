package findgooddaystorobthebank

func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	left, right := make([]int, n), make([]int, n)

	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
		if security[n-i-1] <= security[n-i] {
			right[n-i-1] = right[n-i] + 1
		}
	}
	// i-time>=0 && i+time<=n-1
	for i := time; i <= n-time-1; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return
}
