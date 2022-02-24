package wherewilltheballfall

import (
	"fmt"
	"testing"
)

func TestFindBall(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1},
	}
	fmt.Println(findBall(grid))
}
