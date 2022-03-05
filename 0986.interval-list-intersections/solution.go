package intervallistintersections

func intervalIntersection(firstList [][]int, secondList [][]int) (ans [][]int) {
	m, n := len(firstList), len(secondList)
	i, j := 0, 0
	for i < m && j < n {
		firstLow, firtHigh, secondLow, secondHigh := firstList[i][0], firstList[i][1], secondList[j][0], secondList[j][1]
		// 没有交集
		if firstLow > secondHigh {
			j++
			continue
		}
		if secondLow > firtHigh {
			i++
			continue
		}
		ans = append(ans, []int{max(firstLow, secondLow), min(firtHigh, secondHigh)})
		if firtHigh > secondHigh {
			j++
		} else if firtHigh == secondHigh {
			i++
			j++
		} else {
			i++
		}
	}

	return
}

// 官方简洁版
func intervalIntersection1(A [][]int, B [][]int) (ans [][]int) {
	m, n := len(A), len(B)
	i, j := 0, 0
	for i < m && j < n {
		lo, hi := max(A[i][0], B[j][0]), min(A[i][1], B[j][1])
		// 有交集
		if lo <= hi {
			ans = append(ans, []int{lo, hi})
		}
		// 移除小端
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
