package main

import "fmt"

// struct definition
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // Embedding structs
	// contactInfo can also do this instead -> contactInfo contactInfo
}

func main() {
	// 3 struct declaration options
	//myPerson := person{"Alex", "Anderson"}

	/*
		var myPerson person
		myPerson.firstName = "Alex"
		myPerson.lastName = "Anderson"
		fmt.Printf("%+v", myPerson) // %+v prints out all field names of a struct

		myPerson := person{
			firstName: "Alex",
			lastName:  "Anderson",
		}
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}
	jim.updateName("Jimmy")
	jim.print()

}

// This receiver is pass by reference bc we use a pointer
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// This receiver is pass by value
func (p person) print() {
	fmt.Printf("%+v", p)
}
