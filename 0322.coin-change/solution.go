package coinchange

// 动归,自底向上
func coinChange(coins []int, amount int) int {
	// F(n) = 1+min{F(n-coins[0], ... , F(n-coins[i]))}  (F(n) != -1)
	dp := make([]int, amount+1)
	// dp[0] = 0
	for i := 1; i <= amount; i++ {
		ans := 1<<31 - 1
		for j := range coins {
			if i-coins[j] >= 0 && dp[i-coins[j]] > 0 {
				ans = min(ans, dp[i-coins[j]])
			}
		}
		if ans == (1<<31 - 1) {
			dp[i] = -1
		} else {
			dp[i] = ans + 1
		}

	}

	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 动归+记忆优化搜索,自顶向下
func coinChange1(coins []int, amount int) int {
	memo := make([]int, amount+1)
	// dp 数组全都初始化为特殊值
	for i := range memo {
		memo[i] = -666
	}
	return dp(coins, amount, memo)
}

func dp(coins []int, amount int, memo []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	// 查备忘录，防止重复计算
	if memo[amount] != -666 {
		return memo[amount]
	}

	res := 1<<31 - 1
	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := dp(coins, amount-coin, memo)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		res = min(res, subProblem+1)
	}
	// 把计算结果存入备忘录
	if res == 1<<31-1 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}
