package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// var alex person

	// alex = person{"Alex", "Anderson"}

	// alex = person{firstName: "Alex", lastName: "Anderson"}

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// }

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jimp@gmail.com",
			zipCode: 94000,
		},
	}

	// fmt.Printf("%+v", jim)
	// jimPointer := &jim
	jim.updateName("jimmmm")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
