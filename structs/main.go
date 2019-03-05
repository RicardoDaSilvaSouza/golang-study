package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {

	p := person{
		firstName: "Test1",
		lastName:  "LastNameTest1",
		contactInfo: contactInfo{
			email:   "test1@email.com",
			zipCode: 1234567,
		},
	}

	p.print()
	p.updateName("Test1 Updated")
	p.print()
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
