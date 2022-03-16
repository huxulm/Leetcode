package utf8validation

const mask1, mask2 = 1 << 7, 1<<7 | 1<<6

func validUtf8(data []int) bool {
	n := len(data)
	for i := 0; i < n; {
		if data[i]&mask1 == 0 { // 1字节
			i++
		} else { // data[i] >= 128
			cnt := 0
			for mask := mask1; data[i]&mask != 0; mask >>= 1 {
				cnt++
			}
			if cnt < 2 || cnt > 4 {
				return false
			}
			for j := 1; j < cnt; j++ {
				if i+j >= n || data[i+j]&mask2 != mask1 {
					return false
				}
			}
			i += cnt
		}
	}
	return true
}
