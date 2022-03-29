package maxconsecutiveonesiii

func longestOnes(nums []int, k int) (ans int) {
	n, zeros := len(nums), 0

	l, r := 0, 0
	for r < n {
		if nums[r] == 0 {
			zeros++
		}

		for l < r && zeros > k {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}

		if zeros <= k && r-l+1 > ans {
			ans = r - l + 1
		}

		r++
	}
	return
}
