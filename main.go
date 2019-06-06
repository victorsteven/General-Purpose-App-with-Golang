package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	//declare a variable of type person, third way to declare struct
	var alex person
	alex.firstName = "steven"
	alex.lastName = "Oke"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
