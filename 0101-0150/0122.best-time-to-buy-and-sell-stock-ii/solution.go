package besttimetobuyandsellstockii

func maxProfit(prices []int) (ans int) {
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			ans += prices[i+1] - prices[i]
		}
	}
	return
}
