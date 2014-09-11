package main

import (
	"fmt"

	bins "github.com/SimonRichardson/binoculars/binoculars"
)

var (
	firstName bins.Lens = bins.ObjectLens("FirstName")
	lastName  bins.Lens = bins.ObjectLens("LastName")
	age       bins.Lens = bins.ObjectLens("Age")
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

func (p Person) SetFirstName(v string) Person {
	return firstName.Run(p).Set(v).(Person)
}

func (p Person) SetLastName(v string) Person {
	return lastName.Run(p).Set(v).(Person)
}

func (p Person) SetAge(v uint) Person {
	return age.Run(p).Set(v).(Person)
}

func main() {
	person0 := NewPerson("John", "Doe", 24)
	person1 := person0.SetAge(22)

	fmt.Println(person0, person1)
}
