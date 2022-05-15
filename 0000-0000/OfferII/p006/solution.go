package p006

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	l, r := 0, n-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else { // sum == target
			return []int{l, r}
		}
	}
	return []int{-1, -1}
}
