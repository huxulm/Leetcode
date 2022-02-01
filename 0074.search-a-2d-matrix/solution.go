package searcha2dmatrix

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
