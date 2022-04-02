package arrayofdoubledpairs

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	m := make(map[int]int)
	sort.Ints(arr)

	for i := range arr {
		m[arr[i]]++
	}

	for i := range arr {
		if m[arr[i]] == 0 {
			continue
		}

		if arr[i] < 0 {
			if arr[i]%2 != 0 || m[arr[i]/2] == 0 {
				return false
			} else {
				m[arr[i]]--
				m[arr[i]/2]--
			}
		} else if arr[i] >= 0 {
			if m[arr[i]*2] == 0 {
				return false
			} else {
				m[arr[i]]--
				m[arr[i]*2]--
			}
		}
	}
	return true
}
