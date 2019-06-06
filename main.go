package main

import "fmt"

type conntactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   conntactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	//declare a variable of type person, third way to declare struct
	// var alex person
	// alex.firstName = "steven"
	// alex.lastName = "Oke"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Ade",
		contact: conntactInfo{
			email:   "jim@jim.com",
			zipCode: 17882,
		},
	}

	fmt.Printf("%+v", jim)
}
