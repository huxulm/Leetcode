package stringtointegeratoi

func myAtoi(s string) (res int) {
	// 去除前导空格
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		s = s[i:]
		break
	}

	// + or -
	var sig bool

	for i := 0; i < len(s); i++ {
		if i == 0 && (s[i] == '-' || s[i] == '+') {
			if s[i] == '-' {
				sig = true
			}
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + int(s[i]-'0')
			if !sig {
				if res < 0 || res > 1<<31-1 {
					return 1<<31 - 1
				}
			} else {
				if -res < -1<<31 {
					return -1 << 31
				}
			}
		} else {
			// 不是数字
			break
		}
	}
	if sig {
		res = -res
	}
	return
}
