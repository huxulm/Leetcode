package findtheclosestpalindrome

import (
	"math"
	"strconv"
)

// 1 <= n.length <= 18
// n 只由数字组成
// n 不含前导 0
// n 代表在 [1, 1018 - 1] 范围内的整数
func nearestPalindromic(s string) string {
	n := len(s)
	candidates := []int{int(math.Pow10(n-1)) - 1, int(math.Pow10(n)) + 1}
	prefix, _ := strconv.Atoi(s[:(n+1)/2])
	for _, x := range []int{prefix - 1, prefix, prefix + 1} {
		y := x
		if n&1 == 1 {
			y /= 10
		}
		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candidates = append(candidates, x)
	}
	ans := -1
	selfNum, _ := strconv.Atoi(s)
	for _, c := range candidates {
		if c != selfNum {
			if ans == -1 ||
				abs(c-selfNum) < abs(ans-selfNum) ||
				abs(c-selfNum) == abs(ans-selfNum) && c < ans {
				ans = c
			}
		}
	}

	return strconv.Itoa(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
