package main

import (
	"math"
)

var (
	a int = 6
	b int = 4
)

func RectangleRotation(a, b int) int {
	pt := 0
	radius := math.Pow(float64(a/2), 2) + math.Pow(float64(b/2), 2)
	radius = math.Ceil(math.Pow(float64(radius), 0.5))
	for i := -1 * radius; i < radius+1; i++ {
		for j := -1 * radius; j < radius+1; j++ {
			x := i*math.Cos(-45*(math.Pi/180)) - j*math.Sin(-45*(math.Pi/180))
			y := i*math.Sin(-45*(math.Pi/180)) + j*math.Cos(-45*(math.Pi/180))
			if (-1*float64(a/2) <= x && x <= float64(a/2)) && (-1*float64(b/2) <= y && y <= float64(b/2)) {
				pt += 1
			}

		}
	}
	return pt
	// your code here
}

func main() {

}
