package countnumberofpairswithabsolutedifferencek

func countKDifference(nums []int, k int) int {

	n := len(nums)
	if n < 2 {
		return 0
	}

	ans := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(nums[i]-nums[j]) == k {
				ans++
			}
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
