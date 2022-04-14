package richestcustomerwealth

func maximumWealth(accounts [][]int) (ans int) {
	ans = -1 << 31
	for i := range accounts {
		sum := 0
		for j := range accounts[i] {
			sum += accounts[i][j]
		}
		if sum > ans {
			ans = sum
		}
	}
	return
}
