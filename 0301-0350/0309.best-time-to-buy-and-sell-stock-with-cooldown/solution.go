package besttimetobuyandsellstockwithcooldown

// 持有一支股票，对应的「累计最大收益」记为 f[i][0]；
// 不持有一支股票且处于冷冻期，对应的「累计最大收益」记为 f[i][1]；
// 不持有一支股票且处于非冷冻期，对应的「累计最大收益」记为 f[i][2]；
func maxProfit(prices []int) int {
	n := len(prices)
	f := make([][3]int, n)

	f[0][0], f[0][1], f[0][2] = -prices[0], 0, 0

	for i := 1; i < n; i++ {
		// 当天持有 = max(前一天持有，前一天不持有+当天买)
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		// 当天不持有+冷冻 = (前一天持有+当天天价格)
		f[i][1] = f[i-1][0] + prices[i]
		// 当天不持有+非冷冻 = max(前一天不持有)
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 空间优化
func maxProfit1(prices []int) int {
	n := len(prices)
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		// 当天持有 = max(前一天持有，前一天不持有+当天买)
		newf0 := max(f0, f2-prices[i])
		// 当天不持有+冷冻 = (前一天持有+当天卖的价格)
		newf1 := f0 + prices[i]
		// 当天不持有+非冷冻 = max(前一天不持有(冷冻或不冷冻))
		newf2 := max(f1, f2)
		f0, f1, f2 = newf0, newf1, newf2
	}
	return max(f1, f2)
}
