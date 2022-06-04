package senderwithlargestwordcount

import (
	"strings"
)

func largestWordCount(messages []string, senders []string) (ans string) {
	cnt := map[string]int{}
	for i, msg := range messages {
		n := len(strings.Split(msg, " "))
		cnt[senders[i]] += n
	}

	maxCnt := 0
	for k, v := range cnt {
		if v > maxCnt {
			maxCnt = v
			ans = k
		} else if v == maxCnt && k > ans {
			ans = k
		}
	}
	return
}
