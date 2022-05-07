package openthelock

func openLock(deadends []string, target string) int {
	ds := map[string]bool{}
	for _, v := range deadends {
		ds[v] = true
	}
	vis := map[string]bool{}
	q := []string{"0000"}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, s := range tmp {
			for i, x := range s {
				next := []string{}
				if x-'0' == 0 {
					next = append(next, s[:i]+"9"+s[i+1:])
				} else {
					next = append(next, s[:i]+string(x-1)+s[i+1:])
				}
				next = append(next, s[:i]+string('0'+(x-'0'+1)%10)+s[i+1:])
				for _, v := range next {
					if ds[v] {
						continue
					}
					if v == target {
						return step + 1
					}
					if !vis[v] {
						q = append(q, v)
						vis[v] = true
					}
				}
			}
		}
	}
	return -1
}
