package uniqueemailaddresses

import "strings"

func numUniqueEmails(emails []string) int {
	cnt := map[string]struct{}{}
	for _, addr := range emails {
		at := strings.Index(addr, "@")
		if plus := strings.Index(addr[:at], "+"); plus >= 0 {
			addr = strings.ReplaceAll(addr[:plus], ".", "") + addr[at:]
		} else {
			addr = strings.ReplaceAll(addr[:at], ".", "") + addr[at:]
		}
		cnt[addr] = struct{}{}
	}
	return len(cnt)
}
