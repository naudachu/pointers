package sql

import "pointers/domain/point"

type PointSQL struct {
	ID int
	point.Point
}
