package mergetripletstoformtargettriplet

func mergeTriplets(triplets [][]int, target []int) bool {
	ta, tb, tc := target[0], target[1], target[2]
	var hasA, hasB, hasC bool
	for _, t := range triplets {
		if hasA && hasB && hasC {
			break
		}
		if t[0] <= ta && t[1] <= tb && t[2] <= tc {
			if t[0] == ta {
				hasA = true
			}
			if t[1] == tb {
				hasB = true
			}
			if t[2] == tc {
				hasC = true
			}
		}
	}
	return hasA && hasB && hasC
}
