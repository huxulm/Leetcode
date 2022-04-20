package longestabsolutefilepath

// dir
// ⟶ subdir1
// ⟶ ⟶ file1.ext
// ⟶ ⟶ subsubdir1
// ⟶ subdir2
// ⟶ ⟶ subsubdir2
// ⟶ ⟶ ⟶ file2.ext
func lengthLongestPath(s string) (ans int) {

	n := len(s)
	// 存储当前目录每一级的长度
	level := []int{0}

	var (
		depth   int
		fileLen int
	)
	var isFile bool
	for i := 0; i < n; i++ {
		if s[i] == '\t' {
			depth++
			continue
		} else if s[i] != '\n' && i != n-1 {
			fileLen++
			if s[i] == '.' {
				isFile = true
			}
		} else if s[i] == '\n' || i == n-1 {
			if i == n-1 {
				fileLen++
			}
			level = level[:depth+1]
			if !isFile {
				level = append(level, level[depth]+fileLen+1)
			} else {
				ans = Max(ans, level[depth]+fileLen)
			}

			// reset
			depth, fileLen = 0, 0
			isFile = false
		}

	}
	return
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
