package main

import (
	"example-basics/internal/db"
	"example-basics/internal/model"
	"fmt"
)

func main() {
	database := db.NewDatabase()

	id := database.AddPerson(model.NewPerson("Mehdi", 30))
	fmt.Println(id)

	id = database.AddPerson(model.NewPerson("Pegah", 10))
	fmt.Println(id)

	database.PrintAll()

	ok := database.RemovePerson(1)
	fmt.Println(ok)

	database.PrintAll()

}
