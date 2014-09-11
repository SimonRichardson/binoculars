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

type Phone struct {
	Code   string
	Number string
}

func NewPhone(num string) Phone {
	return Phone{
		Code:   "+44",
		Number: num[1:],
	}
}

type Person struct {
	FirstName string
	LastName  string
	Age       uint
	Phone     Phone
}

func NewPerson(firstName, lastName string, age uint, phone string) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Phone:     NewPhone(phone),
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
	// Flat changes!
	person0 := NewPerson("John", "Doe", 24, "07123456789")
	person1 := person0.SetAge(22)

	fmt.Println(person0, person1)

	// Nested changes!
	x := bins.ObjectLens("Phone")
	y := bins.ObjectLens("Number")

	fmt.Println(person0, x.AndThen(y).Run(person0).Set("7987654321"))

	// Multiple changes!
	fmt.Println(person0, firstName.And(lastName).Run(person0).Set([]bins.Any{"A", "B"}))
}
