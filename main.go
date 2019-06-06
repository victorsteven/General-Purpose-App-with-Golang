package main

import "fmt"

type conntactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	conntactInfo
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
		conntactInfo: conntactInfo{
			email:   "jim@jim.com",
			zipCode: 17882,
		},
	}

	// fmt.Printf("%+v", jim)
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

//this function take the person receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
