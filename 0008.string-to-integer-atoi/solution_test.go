package stringtointegeratoi

import (
	"fmt"
	"testing"
)

func TestAtoi(t *testing.T) {
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-9223372036854775808"))
}
