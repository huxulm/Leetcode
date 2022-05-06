package countsubmatriceswithallones

import "container/list"

// 方法一: 枚举
func numSubmat(mat [][]int) (ans int) {
	m, n := len(mat), len(mat[0])
	rows := make([][]int, m)
	for i := range rows {
		rows[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				rows[i][j] = mat[i][j]
				continue
			}
			if mat[i][j] == 0 {
				rows[i][j] = 0
			} else {
				rows[i][j] = rows[i][j-1] + 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			col := rows[i][j]
			for k := i; k >= 0 && col > 0; k-- {
				col = min(col, rows[k][j])
				ans += col
			}
		}
	}

	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 单调栈
func numSubmat1(mat [][]int) (ans int) {
	type pair struct {
		first, second int
	}
	n, m := len(mat), len(mat[0])

	row := make([][]int, n)
	for i := range row {
		row[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				row[i][j] = mat[i][j]
			} else if mat[i][j] == 1 {
				row[i][j] = row[i][j-1] + 1
			} else {
				row[i][j] = 0
			}
		}
	}
	for j := 0; j < m; j++ {
		i := 0
		Q := list.New()
		sum := 0
		for i <= n-1 {
			height := 1
			for Q.Len() > 0 && Q.Back().Value.(pair).first > row[i][j] {
				// 弹出的时候要减去多于的答案
				sum -= Q.Back().Value.(pair).second * (Q.Back().Value.(pair).first - row[i][j])
				height += Q.Back().Value.(pair).second
				Q.Remove(Q.Back())
			}
			sum += row[i][j]
			ans += sum
			Q.PushBack(pair{row[i][j], height})
			i++
		}
	}
	return
}

// 参考: https://leetcode.com/problems/count-submatrices-with-all-ones/discuss/720265/Java-Detailed-Explanation-From-O(MNM)-to-O(MN)-by-using-Stack
func numSubmat2(mat [][]int) (ans int) {
	m := len(mat[0])
	heights := make([]int, m)
	for _, row := range mat {
		sum := make([]int, m)
		type pair struct{ h, j int }
		stack := []pair{{-1, -1}}
		for j, v := range row {
			if v == 0 {
				heights[j] = 0
			} else {
				heights[j]++
			}
			h := heights[j]
			for {
				if top := stack[len(stack)-1]; top.h < h {
					if pre := top.j; pre < 0 {
						sum[j] = (j + 1) * h
					} else {
						sum[j] = sum[pre] + (j-pre)*h
					}
					ans += sum[j]
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{h, j})
		}
	}
	return
}
