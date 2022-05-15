package p003

import "math/bits"

func countBits(n int) (ans []int) {
	for i := 0; i <= n; i++ {
		cnt := 0
		for mask := i; mask > 0; mask >>= 1 {
			if mask&1 == 1 {
				cnt++
			}
		}
		ans = append(ans, cnt)
	}
	return
}

func countBits1(n int) (ans []int) {
	for i := 0; i <= n; i++ {
		ans = append(ans, bits.OnesCount(uint(i)))
	}
	return
}
