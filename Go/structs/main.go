package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	budi := person{
		firstName: "Budi",
		lastName:  "Sudarsono",
		contact: contactInfo{
			email:   "budi@gmail.com",
			zipCode: 123412,
		},
	}

	budi.print()

	budiPointer := &budi
	budiPointer.updateName("badu")
	budi.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
