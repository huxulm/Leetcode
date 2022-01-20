package containerwithmostwater

func maxArea(height []int) (area int) {
	left, right := 0, len(height)-1

	area = 0

	for left < right {
		if height[left] < height[right] {
			area = max(area, (right-left)*height[left])
			left++
		} else {
			area = max(area, (right-left)*height[right])
			right--
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
