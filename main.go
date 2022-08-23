package main

import (
	"pointers/app"
	"pointers/sql"
)

func main() {
	store := initDB("postgres://postgres:helpme123@localhost:5432/postgres")
	app.Work(&store)
}

func initDB(str string) sql.Store {
	return sql.NewStore(str)
}
