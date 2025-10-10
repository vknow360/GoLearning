package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
}

type Address struct {
	city    string
	country string
}

func main() {

	p := Person{
		"John", "Doe", 30,
		Address{
			"ABC", "xyz",
		},
	}

	p1 := Person{
		firstName: "Jane",
		age:       25,
	}

	fmt.Println(p.firstName)
	fmt.Println(p1.fullName())
	fmt.Println(p1.age)
	p1.incrementAge()
	fmt.Println(p1.age)

	user := struct {
		username string
		email    string
	}{
		"user123",
		"abc@example.com",
	}
	fmt.Println(user)

}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAge() { // Pointer receiver
	p.age++
}
