package numberofsmoothdescentperiodsofastock

func getDescentPeriods(prices []int) (ans int64) {
	n := len(prices)

	cnt := 0

	for i := 0; i < n; i++ {
		ans++
		if i < n-1 && prices[i] == prices[i+1]+1 {
			cnt++
			ans += int64(cnt)
		} else {
			cnt = 0
		}
	}
	return
}
