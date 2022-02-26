package besttimetobuyandsellstockwithtransactionfee

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	// f[i][0] 第i天结束无股票最大利润
	// f[i][1] 第i天结束有股票最大利润
	f := make([][2]int, n)
	f[0][0], f[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1]+prices[i]-fee)
		f[i][1] = max(f[i-1][0]-prices[i], f[i-1][1])
	}

	return f[n-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// DP压缩
func maxProfit1(prices []int, fee int) int {
	n := len(prices)
	// f0 第i天结束无股票最大利润
	// f1 第i天结束有股票最大利润
	f0, f1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		f0, f1 = max(f0, f1+prices[i]-fee), max(f0-prices[i], f1)
	}

	return f0
}
