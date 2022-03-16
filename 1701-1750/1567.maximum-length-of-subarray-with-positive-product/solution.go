package maximumlengthofsubarraywithpositiveproduct

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getMaxLen(nums []int) (ans int) {
	// f[i][0], f[i][1] 以第i个数结尾的乘积为正,或负的最长子数组长度
	n := len(nums)
	pos, neg := make([]int, n), make([]int, n)

	if nums[0] > 0 {
		pos[0] = 1
	} else if nums[0] < 0 {
		neg[0] = 1
	}

	ans = pos[0]

	for i := 1; i < n; i++ {
		if nums[i] == 0 {
			pos[i], neg[i] = 0, 0
		}
		if nums[i] > 0 {
			pos[i] = pos[i-1] + 1
			if neg[i-1] > 0 {
				neg[i] = neg[i-1] + 1
			} else {
				neg[i] = 0
			}
		}
		if nums[i] < 0 {
			neg[i] = pos[i-1] + 1
			if neg[i-1] > 0 {
				pos[i] = neg[i-1] + 1
			} else {
				pos[i] = 0
			}
		}

		ans = max(ans, pos[i])
	}

	return
}

// 贪心
func getMaxLen1(nums []int) (ans int) {
	n := len(nums)
	pos, neg, first := 0, 0, -1
	for i := 0; i < n; i++ {
		if nums[i] == 0 { // 重置
			pos, neg, first = 0, 0, -1
		} else if nums[i] > 0 {
			pos++
		} else {
			if first == -1 {
				first = i
			}
			neg++
		}
		// 负数个数为偶数
		if neg%2 == 0 {
			ans = max(ans, neg+pos)
		} else { // 负数个数为奇数，贪心选择：第一个负数后的位置到当前位置的数组乘积
			ans = max(ans, i-first)
		}
	}
	return
}
