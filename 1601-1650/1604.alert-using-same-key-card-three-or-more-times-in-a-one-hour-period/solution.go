package alertusingsamekeycardthreeormoretimesinaonehourperiod

import (
	"sort"
	"strconv"
)

func timeToMiniutes(t string) int {
	m, _ := strconv.Atoi(t[:2])
	s, _ := strconv.Atoi(t[3:])
	return m*60 + s
}

func alertNames(kn []string, kt []string) (ans []string) {
	m := map[string][]int{}

	for i := range kn {
		m[kn[i]] = append(m[kn[i]], timeToMiniutes(kt[i]))
	}

	for k, T := range m {
		if len(T) < 3 {
			continue
		}
		sort.Ints(T)
		for i := 0; i < len(T)-2; i++ {
			if T[i+2]-T[i] <= 60 {
				ans = append(ans, k)
				break
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })
	return
}
