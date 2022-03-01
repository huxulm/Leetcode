package trappingrainwater

func trap(height []int) (ans int) {
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)

	leftMax[0], rightMax[n-1] = height[0], height[n-1]

	l_max := leftMax[0]
	for i := 1; i < n; i++ {
		if height[i] > l_max {
			l_max = height[i]
		}
		leftMax[i] = l_max
	}

	r_max := rightMax[n-1]
	for i := n - 2; i >= 0; i-- {
		if height[i] > r_max {
			r_max = height[i]
		}
		rightMax[i] = r_max
	}

	for i := range height {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return
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

func trap1(height []int) (ans int) {
	n := len(height)
	leftMax, rightMax := 0, 0
	l, r := 0, n-1

	for l < r {
		leftMax, rightMax = max(leftMax, height[l]), max(rightMax, height[r])
		if height[l] < height[r] {
			ans += leftMax - height[l]
		} else {
			ans += rightMax - height[r]
		}
	}
	return
}
