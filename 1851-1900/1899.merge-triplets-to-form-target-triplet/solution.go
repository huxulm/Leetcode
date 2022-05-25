package mergetripletstoformtargettriplet

import (
	"sort"
)

func mergeTriplets(triplets [][]int, target []int) bool {
	ta, tb, tc := target[0], target[1], target[2]

	tmp := [][]int{}
	for _, t := range triplets {
		if t[0] <= ta && t[1] <= tb && t[2] <= tc {
			tmp = append(tmp, t)
		}
	}

	triplets = tmp
	n := len(triplets)
	if n == 0 {
		return false
	}

	sort.Slice(triplets, func(i, j int) bool { return triplets[i][0] < triplets[j][0] })

	if triplets[n-1][0] != ta {
		return false
	}

	sort.Slice(triplets, func(i, j int) bool { return triplets[i][1] < triplets[j][1] })

	if triplets[n-1][1] != tb {
		return false
	}

	sort.Slice(triplets, func(i, j int) bool { return triplets[i][2] < triplets[j][2] })

	if triplets[n-1][2] != tc {
		return false
	}

	return true
}
