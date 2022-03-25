package factorialtrailingzeroes

// [2, 5] 的对数
func trailingZeroes(n int) (ans int) {

	for i := 5; i <= n; i++ {
		for i%5 == 0 {
			ans++
			i /= 5
		}
	}

	return
}
