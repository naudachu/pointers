package point

import (
	"math"
)

type Point struct {
	X, Y int
}

type PointRepository interface {
	SearchPointByID(int) Point
}

func DistanceBetween(i, j Point) float64 {
	dx := float64((j.X - i.X))
	dy := float64((j.Y - i.Y))
	return math.Sqrt(dx*dx + dy*dy)
}
