package countnumberswithuniquedigits

func countNumbersWithUniqueDigits(n int) (ans int) {
	ans = 1
	cur := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			cur = 9
			ans += cur
			continue
		}
		cur *= (10 - i)
		ans += cur
	}
	return
}
