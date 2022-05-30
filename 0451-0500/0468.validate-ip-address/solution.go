package validateipaddress

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if isIpv4(queryIP) {
		return "IPv4"
	} else if isIpv6(queryIP) {
		return "IPv6"
	} else {
		return "Neither"
	}
}

func isIpv4(ip string) bool {
	if len(ip) > 15 || len(ip) < 7 {
		return false
	}
	hasDot := strings.Index(ip, ".") >= 0
	if !hasDot {
		return false
	}
	ss := strings.Split(ip, ".")
	if len(ss) != 4 {
		return false
	}

	for i := range ss {
		l := len(ss[i])
		val := 0
		for _, ch := range ss[i] {
			if ch < '0' || ch > '9' {
				return false
			}
			val = val*10 + int(ch-'0')
		}
		// 含前导 0
		if val > 255 || l == 3 && val < 100 || l == 2 && val < 10 {
			return false
		}
	}
	return true
}

func isIpv6(ip string) bool {
	if len(ip) > 39 || len(ip) < 15 {
		return false
	}
	hasDot := strings.Index(ip, ":") >= 0
	if !hasDot {
		return false
	}
	ss := strings.Split(ip, ":")
	if len(ss) != 8 {
		return false
	}
	for i := range ss {
		if len(ss[i]) > 4 {
			return false
		}
		for _, ch := range ss[i] {
			if ch >= '0' && ch <= '9' || ch >= 'a' && ch <= 'f' || ch >= 'A' && ch <= 'F' {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// LC https://leetcode.cn/problems/validate-ip-address/solution/yan-zheng-ipdi-zhi-by-leetcode-solution-kge5/
func validIPAddress1(queryIP string) string {
	if sp := strings.Split(queryIP, "."); len(sp) == 4 {
		for _, s := range sp {
			if len(s) > 1 && s[0] == '0' {
				return "Neither"
			}
			if v, err := strconv.Atoi(s); err != nil || v > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	if sp := strings.Split(queryIP, ":"); len(sp) == 8 {
		for _, s := range sp {
			if len(s) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseUint(s, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
