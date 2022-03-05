package countthehiddensequences

func numberOfArrays(diff []int, lower int, upper int) int {
	x, y, sum := 0, 0, 0
	for i := range diff {
		sum += diff[i]
		x, y = min(x, sum), max(y, sum)
		if y-x > upper-lower {
			return 0
		}
	}
	return upper - lower - (y - x) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
