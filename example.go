package main

import (
	"fmt"

	"github.com/SimonRichardson/binoculars/binoculars"
)

type Person struct {
	FirstName string
	LastName  string
	Age       uint
}

func NewPerson(firstName, lastName string, age uint) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func PersonAge(p Person, age uint) Person {
	return binoculars.ObjectLens("Age").Run(p).Set(age).(Person)
}

func main() {
	person0 := NewPerson("Timmy", "Rogers", 24)
	person1 := PersonAge(person0, 22)

	fmt.Println(person0, person1)
}
