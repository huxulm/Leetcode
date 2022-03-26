package baseballgame

import "strconv"

func calPoints(ops []string) (ans int) {
	score := []int{}
	for i := range ops {
		if ops[i] == "C" {
			score = score[:len(score)-1]
		} else if ops[i] == "D" {
			score = append(score, score[len(score)-1]*2)
		} else if ops[i] == "+" {
			score = append(score, score[len(score)-1]+score[len(score)-2])
		} else {
			s, _ := strconv.Atoi(ops[i])
			score = append(score, s)
		}
	}

	for _, s := range score {
		ans += s
	}

	return
}
