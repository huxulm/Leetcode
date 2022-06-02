package matchstickstosquare

import "sort"

// 暴力: 二进制枚举
func makesquare(matchsticks []int) (ans bool) {
	n := len(matchsticks)
	if n < 4 {
		return false
	}
	sum := 0
	for i := range matchsticks {
		sum += matchsticks[i]
	}

	if sum%4 != 0 {
		return false
	}

	side := sum / 4
	var s = make([]int, 1<<n)

	// n <= 15
	for mask := 1; mask < 1<<n; mask++ {
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				s[mask] += matchsticks[i]
			}
		}
	}

	for a := 1; a < 1<<n; a++ {
		if s[a] != side {
			continue
		}
		for x, b := (1<<n)-1^a, (1<<n)-1^a; b > 0; b = (b - 1) & x {
			if s[b] != side {
				continue
			}
			for y, c := x^b, x^b; c > 0; c = (c - 1) & y {
				if s[c] != side {
					continue
				}
				// 第四个边可省略，前三个边长度等于side，那第四个边一定满足
				/*int d = y ^ c;
				if (s[d] != side) continue;
				return true;*/
				return true
			}
		}
	}

	return
}

// DFS + 剪枝
func makesquare1(matchsticks []int) (ans bool) {
	n := len(matchsticks)
	sum := 0
	for i := range matchsticks {
		sum += matchsticks[i]
	}
	if n < 4 || sum%4 != 0 {
		return false
	}

	side := sum / 4

	var edges [4]int

	// 排序剪枝
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == n {
			return true
		}
		for i := range edges {
			if edges[i] == side || edges[i]+matchsticks[idx] > side || (i > 0 && edges[i] == edges[i-1]) {
				continue
			}
			edges[i] += matchsticks[idx]
			if dfs(idx + 1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}

	return dfs(0)
}
