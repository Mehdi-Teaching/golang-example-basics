package db

import "example-basics/internal/model"

type Database struct {
	people map[uint]model.Person
}

func NewDatabase() Database {
	return Database{
		people: make(map[uint]model.Person),
	}
}
