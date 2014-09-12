package main

import (
	"fmt"

	bins "github.com/SimonRichardson/binoculars/binoculars"
)

var (
	propFirstName bins.Property = "FirstName"
	propLastName  bins.Property = "LastName"
	propAge       bins.Property = "Age"
	propPhone     bins.Property = "Phone"
	propNumber    bins.Property = "Number"

	lenses map[bins.Property]bins.Lens = bins.ObjectLenses([]bins.Property{propFirstName, propLastName, propAge})
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
	return lenses[propFirstName].Run(p).Set(v).(Person)
}

func (p Person) SetLastName(v string) Person {
	return lenses[propLastName].Run(p).Set(v).(Person)
}

func (p Person) SetAge(v uint) Person {
	return lenses[propAge].Run(p).Set(v).(Person)
}

func main() {
	// Flat changes!
	person0 := NewPerson("John", "Doe", 24, "07123456789")
	person1 := person0.SetAge(22)

	fmt.Println(person0, person1)

	// Nested changes!
	x := bins.ObjectLens(propPhone)
	y := bins.ObjectLens(propNumber)

	fmt.Println(person0, x.AndThen(y).Run(person0).Set("7987654321"))

	// Partial changes!

	a := bins.ObjectPartialLens(propFirstName)
	b := bins.ObjectPartialLens("X")

	change := func(x bins.Any) bins.Any {
		return x.(bins.Store).Set("xxxxx")
	}

	fmt.Println(person0, a.Run(person0).Map(change), b.Run(person0).Map(change))

	// Functor changes!
	firstName := func(a bins.Any) bins.Any {
		return a.(Person).FirstName
	}

	fmt.Println(person0, lenses[propFirstName].Run(person0).Map(firstName).Extract())
}
