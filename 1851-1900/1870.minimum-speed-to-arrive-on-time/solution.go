package minimumspeedtoarriveontime

import "sort"

func minSpeedOnTime(dist []int, hour float64) int {

	n := len(dist)
	if hour <= float64(n-1) {
		return -1
	}

	var check = func(s int) bool {
		s++
		// 前 n - 1 趟时间
		var tot int
		for _, d := range dist[:n-1] {
			tot += (d + s - 1) / s
			if float64(tot) >= hour {
				return false
			}
		}
		return float64(tot)+float64(dist[n-1])/float64(s) <= hour
	}

	return sort.Search(1e7, check) - 1
}
