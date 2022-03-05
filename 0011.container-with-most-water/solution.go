package containerwithmostwater

func maxArea(height []int) (area int) {
	lo, hi := 0, len(height)-1

	for lo < hi {
		if height[lo] < height[hi] {
			area = max(area, height[lo]*(hi-lo))
			lo++
		} else {
			area = max(area, height[hi]*(hi-lo))
			hi--
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
