package wordsearch

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	// board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	// word := "ASADFBCCEESE"
	board := [][]byte{{'a'}}
	word := "a"
	fmt.Println(exist(board, word))
}
