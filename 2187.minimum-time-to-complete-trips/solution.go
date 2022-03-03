package minimumtimetocompletetrips

import "sort"

func minimumTime(time []int, totalTrips int) int64 {
	return int64(sort.Search(totalTrips*1e7, func(x int) bool {
		tot := 0
		for i := range time {
			tot += x / time[i]
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}))
}
