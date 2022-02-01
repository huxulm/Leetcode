package searcha2dmatrix

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lo, hi := 0, m-1

	for lo <= hi {
		mid := (lo + hi) >> 1
		mid_min, mid_max := matrix[mid][0], matrix[mid][n-1]
		if lo == hi || (target >= mid_min && target <= mid_max) {
			return binarySearch(matrix[mid], target)
		}
		if target < mid_min {
			hi = mid - 1
		} else if target > mid_max {
			lo = mid + 1
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

// LC official
func searchMatrix1(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
