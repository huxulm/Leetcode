package smallestrotationwithhighestscore

// 输入：nums = [2,3,1,4,0]
// 输出：3
// 解释：
// 下面列出了每个 k 的得分：
// k = 0,  nums = [2,3,1,4,0],    score 2
// k = 1,  nums = [3,1,4,0,2],    score 3
// k = 2,  nums = [1,4,0,2,3],    score 3
// k = 3,  nums = [4,0,2,3,1],    score 4
// k = 4,  nums = [0,2,3,1,4],    score 3
func bestRotation(nums []int) int {
	n := len(nums)
	diffs := make([]int, n)
	for i, num := range nums {
		low := (i + 1) % n
		high := (i - num + n + 1) % n
		diffs[low]++
		diffs[high]--
		if low >= high {
			diffs[0]++
		}
	}
	score, maxScore, idx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore, idx = score, i
		}
	}
	return idx
}
