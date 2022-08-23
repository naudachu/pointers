package main

import (
	"fmt"
	"pointers/domain/point"
	"pointers/sql"
)

func main() {
	store := initDB("postgres://postgres:helpme123@localhost:5432/postgres")
	//store.Conn.Acquire(context.Background())
	//var p1 point.Point
	p1 := store.SearchPointByID(1)
	p2 := store.SearchPointByID(2)

	fmt.Print(point.DistanceBetween(p1, p2))
}

func initDB(str string) sql.Store {
	return sql.NewStore(str)
}
