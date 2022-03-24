package main

import "fmt"

type contactInfo struct {
	email   string `json:"email"`
	zipcode int    `json:"zipcode"`
}

type person struct {
	firstName   string `json:"first_name"`
	lastName    string `json:"last_name"`
	age         int    `json:"age"`
	contactInfo `json:"contact_info"`
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	nathalie := person{firstName: "Nathalie", lastName: "Croeghs", age: 32}
	fmt.Printf("First Name: %v\n", nathalie.firstName)
	nathalie.age = 33
	fmt.Println(nathalie)

	jim := person{
		firstName: "Jim",
		lastName:  "Dupont",
		age:       36,
		contactInfo: contactInfo{
			email:   "jim@no.no",
			zipcode: 12540,
		},
	}

	jim.updateName("Jimmy")

	jim.print()

}
