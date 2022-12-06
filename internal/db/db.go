package db

import (
	"example-basics/internal/model"
	"fmt"
)

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

func (d *Database) AddPerson(person model.Person) uint {
	d.people[d.counter] = person

	id := d.counter

	d.counter++

	return id
}

func (d *Database) RemovePerson(id uint) bool {
	_, ok := d.people[id]
	if !ok {
		return false
	}

	delete(d.people, id)

	return true
}

func (d *Database) GetAll() []model.Person {
	result := make([]model.Person, 0)

	for _, person := range d.people {
		result = append(result, person)
	}

	return result
}

func (d *Database) PrintAll() {
	fmt.Printf("%+v\n", d.GetAll())
}
