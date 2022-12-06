package db

import "example-basics/internal/model"

type Database struct {
	people  map[uint]model.Person
	counter uint
}

func NewDatabase() Database {
	return Database{
		people:  make(map[uint]model.Person),
		counter: 0,
	}
}
