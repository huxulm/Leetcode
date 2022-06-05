package generaterandompointinacircle

import "math/rand"

type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius: radius, xCenter: x_center, yCenter: y_center}
}

func (s *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if x*x+y*y < 1 {
			return []float64{s.xCenter + x*s.radius, s.yCenter + y*s.radius}
		}
	}
}
