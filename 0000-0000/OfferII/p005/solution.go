package p005

func maxProduct(words []string) (ans int) {
	n := len(words)
	keys := make([]int, n)
	for i, s := range words {
		k := 0
		for _, ch := range s {
			k |= 1 << (ch - 'a')
		}
		keys[i] = k
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if keys[i]&keys[j] > 0 {
				continue
			}
			if v := len(words[i]) * len(words[j]); v > ans {
				ans = v
			}
		}
	}
	return
}
