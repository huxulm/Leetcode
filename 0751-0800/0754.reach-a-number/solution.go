package reachanumber

import "math"

func reachNumber(target int) (ans int) {
	target = int(math.Abs(float64(target)))
	k := 0
	for target > 0 {
		k++
		target -= k
	}

	// delta 为偶数
	if target&1 == 0 {
		ans = k
	} else {
		// delta 为奇数, 考虑 k + 1 和 k + 2
		ans = k + 1 + k%2
	}
	return
}
