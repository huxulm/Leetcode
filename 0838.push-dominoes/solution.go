package pushdominoes

func pushDominoes(dominoes string) string {
	dominoes = "L" + dominoes + "R"
	l := 0
	res := ""

	for r := 1; r < len(dominoes); r++ {
		if dominoes[r] == '.' {
			continue
		}
		if l != 0 {
			res += string(dominoes[l])
		}
		mid := r - l - 1
		if dominoes[l] == dominoes[r] {
			for i := 0; i < mid; i++ {
				res += string(dominoes[l])
			}
		} else if dominoes[l] == 'L' && dominoes[r] == 'R' {
			for i := 0; i < mid; i++ {
				res += string('.')
			}
		} else {
			for i := 0; i < mid/2; i++ {
				res += string('R')
			}
			if mid%2 == 1 {
				res += string('.')
			}
			for i := 0; i < mid/2; i++ {
				res += string('L')
			}
		}
		l = r
	}

	return res
}
