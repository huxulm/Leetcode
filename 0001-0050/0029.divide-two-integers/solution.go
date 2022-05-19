package dividetwointegers

import (
	"math"
)

// LC 官方
func divide(dividend int, divisor int) int {
	// 被除数为 0
	if dividend == 0 {
		return 0
	}

	if dividend == math.MinInt32 { // 被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}

	if divisor == math.MinInt32 { // 考虑除数为最小值情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}

	// 一般情况, 使用二分查找
	// 将所有的正数取相反数, 这样只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}

	ans := 0
	l, r := 1, math.MaxInt32
	for l <= r {
		mid := int(uint(l+r) >> 1)
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 {
				break
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if rev {
		ans = -ans
	}
	return ans
}

// 快速乘
// x 和 y 是负数，z 是正数
// 判断 z * y >= x 是否成立
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // 不使用除法
		if z&1 > 0 {
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			// 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func divide1(dividend int, divisor int) int {
	x, y := dividend, divisor
	isNeg := (x > 0 && y < 0) || (x < 0 && y > 0)

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	// 二分右边界 答案区间: [0, x+1)
	l, r := 0, x+1
	ans := 0

	for l < r {
		mid := int(uint(l+r) >> 1)
		if mul(y, mid) > x { // 找到最小的 满足 mul(a, x) = true 的 x
			r = mid
		} else {
			l = mid + 1
		}
	}

	ans = l - 1

	// STL 一行流
	// ans = sort.Search(x+1, func(i int) bool { return mul(y, i) > x }) - 1

	if ans == -1 {
		ans = 0
	}
	if isNeg {
		ans = -ans
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		ans = math.MaxInt32
	}
	return ans
}

func mul(a, b int) (ans int) {
	for b > 0 {
		if b&1 == 1 {
			ans += a
		}
		a <<= 1
		b >>= 1
	}
	return
}
