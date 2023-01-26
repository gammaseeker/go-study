package main

import "fmt"

// slice, maps, channels, pointers, functions are passed in by reference
// structs, int, float, string, bool are passed by value unless you use a pointer
// Slices are like lists from python or arrayList from Java, an array is the traditional array

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
	/*
		jimPointer := &jim // &variable provides memory address that variable is pointing at
		// *pointer gives value that the mem addr is pointing at
		jimPointer.updateName("Jimmy")
	*/

	// Golang also lets you use the value as a receiver for a pointer function
	jim.updateName("Jimmy")
	jim.print()

}

// This receiver is pass by reference bc we use a pointer. This type of receiver function can be used
// by the pointer or variable
func (p *person) updateName(newFirstName string) {
	// *type means we are working with a pointer to that type
	(*p).firstName = newFirstName
	// *p it means we want to manipulate the value the pointer is referencing
}

// This receiver is pass by value
func (p person) print() {
	fmt.Printf("%+v", p)
}
