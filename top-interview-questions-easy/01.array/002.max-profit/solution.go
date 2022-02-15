package maxprofit

func maxProfit(prices []int) int {
	n := len(prices)

	if n <= 1 {
		return 0
	}

	ans := 0

	for i := 0; i < n-1; i++ {
		if prices[i] < prices[i+1] {
			ans += prices[i+1] - prices[i]
		}
	}

	return ans
}
