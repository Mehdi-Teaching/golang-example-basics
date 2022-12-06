package model

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}
