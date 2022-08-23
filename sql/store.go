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
	var tempPoint point.Point
	tempPoint.X = s.X
	tempPoint.Y = s.Y
	return s
}
