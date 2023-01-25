package main

// struct definition
type person struct {
	firstName string
	lastName  string
}

func main() {
	// 3 struct declaration options
	//myPerson := person{"Alex", "Anderson"}

	/*
		var myPerson person
		myPerson.firstName = "Alex"
		myPerson.lastName = "Anderson"
		fmt.Printf("%+v", myPerson) // %+v prints out all field names of a struct
	*/

	myPerson := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

}
