package app

import (
	"fmt"
	"pointers/domain/point"
	"pointers/sql"
)

func Work(store *sql.Store) {
	p1 := store.SearchPointByID(1)
	p2 := store.SearchPointByID(3)

	list := store.SearhAllPoints()

	fmt.Print(point.DistanceBetween(p2, p1))
	fmt.Print(list)
}
