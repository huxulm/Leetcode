package selfdividingnumbers

func selfDividingNumbers(left int, right int) (ans []int) {
	for x := left; x <= right; x++ {
		if isDividing(x) {
			ans = append(ans, x)
		}
	}
	return
}

func isDividing(n int) bool {
	for x := n; x > 0; x /= 10 {
		if v := x % 10; v == 0 || n%v != 0 {
			return false
		}
	}
	return true
}
