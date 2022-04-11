package executionofallsuffixinstructionsstayinginagrid

// 模拟
var dirs = [][2]int{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}

func executeInstructions1(n int, startPos []int, s string) []int {
	ans := make([]int, len(s))
	for i := range s {
		ans[i] = len(s) - i
		x, y := startPos[0], startPos[1]
		for j, ch := range s[i:] {
			x += dirs[ch][0]
			y += dirs[ch][1]
			if x < 0 || x >= n || y < 0 || y >= n {
				ans[i] = j
				break
			}
		}
	}
	return ans
}

// 前缀和
func executeInstructions(n int, startPos []int, s string) (ans []int) {
	m := len(s)

	var offset = make([][2]int, m+1)

	for i := 1; i <= m; i++ {
		if s[i-1] == 'U' { // row
			offset[i][0] = offset[i-1][0] - 1
			offset[i][1] = offset[i-1][1]
		} else if s[i-1] == 'D' { // row
			offset[i][0] = offset[i-1][0] + 1
			offset[i][1] = offset[i-1][1]
		} else if s[i-1] == 'L' { // col
			offset[i][0] = offset[i-1][0]
			offset[i][1] = offset[i-1][1] - 1
		} else if s[i-1] == 'R' { // col
			offset[i][0] = offset[i-1][0]
			offset[i][1] = offset[i-1][1] + 1
		}
	}

	ans = make([]int, m)

	for i := range s {
		cnt := 0

		for j := i + 1; j <= m; j++ {
			x := startPos[0] + offset[j][0] - offset[i][0]
			y := startPos[1] + offset[j][1] - offset[i][1]
			if (x >= 0 && x < n) && (y >= 0 && y < n) {
				cnt++
				continue
			} else {
				break
			}
		}
		ans[i] = cnt
	}

	return
}
