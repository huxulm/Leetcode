package main

import (
	"bytes"
	"fmt"
)

func addBinary(a string, b string) string {
	sb := bytes.Buffer{}
	p, q := len(a)-1, len(b)-1

	var flag byte
	for p >= 0 || q >= 0 {
		v := flag
		if p >= 0 {
			v += a[p] - '0'
			p--
		}
		if q >= 0 {
			v += b[q] - '0'
			q--
		}
		// next flag
		flag = v >> 1
		sb.WriteByte(v%2 + '0')
	}
	if flag > 0 {
		sb.WriteByte(flag + '0')
	}
	res := sb.Bytes()
	n := len(res)
	for i := 0; i < n>>1; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return string(res)
}

func main() {
	fmt.Println(addBinary("10", "11"))
}
