package sql

import (
	"context"
	"log"
	"pointers/domain/point"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	Conn *pgxpool.Pool
	PointSQL
}

func NewStore(str string) (s Store) {
	var err error

	s.Conn, err = pgxpool.Connect(context.TODO(), str)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func (s Store) SearchPointByID(id int) point.Point {
	s = s.searchPointByID(id)

	return point.Point{
		X: s.X,
		Y: s.Y,
	}
}
func (s Store) searchPointByID(id int) Store {
	s.Conn.QueryRow(context.Background(),
		"SELECT * FROM point WHERE id = $1",
		id).Scan(&s.X, &s.Y, &s.ID)
	return s
}

func (s Store) SearhAllPoints() []point.Point {
	//var list []point.Point
	stores := s.searchAllPoints()
	var list []point.Point
	for _, val := range stores {
		list = append(list, point.Point{
			X: val.X,
			Y: val.Y,
		})

	}
	return list
}

func (s Store) searchAllPoints() []Store {
	var stores []Store
	rows, err := s.Conn.Query(context.Background(),
		"SELECT * FROM point")
	if err != nil {
		log.Print(err)
	}
	//.Scan(&s.X, &s.Y, &s.ID)
	for rows.Next() {
		err := rows.Scan(&s.X, &s.Y, &s.ID)
		if err != nil {
			log.Print(err)
		}
		stores = append(stores, s)
	}
	return stores
}
