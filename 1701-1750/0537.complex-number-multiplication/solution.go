package complexnumbermultiplication

import (
	"fmt"
	"strconv"
	"strings"
)

// 输入：num1 = "1+1i", num2 = "1+1i"
// 输出："0+2i"
func complexNumberMultiply(num1 string, num2 string) string {
	a, b := 0, 0
	a1, b1 := parse(num1)
	a2, b2 := parse(num2)
	a, b = a1*a2-b1*b2, a1*b2+b1*a2
	return fmt.Sprintf("%d+%di", a, b)
}

func parse(str string) (a, b int) {
	res := strings.Split(str[:len(str)-1], "+")
	a, _ = strconv.Atoi(res[0])
	b, _ = strconv.Atoi(res[1])
	return
}
