package app

import (
	"fmt"
	"pointers/domain/point"
	"pointers/sql"
)

func Work(db *sql.Store) {
	p1 := db.SearchPointByID(1)
	p2 := db.SearchPointByID(3)
	fmt.Printf("Distance between %d:%d and %d:%d = %f\n", p1.X, p1.Y, p2.X, p2.Y, point.DistanceBetween(p2, p1))

	list := db.SearhAllPoints()

	fmt.Print(list)
}
