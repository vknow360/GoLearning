package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person
	employeeID string
	salary     float64
}

func (p person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

func (e Employee) greet() {
	fmt.Printf("Hello, my name is %s, I am %d years old, and my employee ID is %s.\n", e.name, e.age, e.employeeID)
}

func main() {
	emp := Employee{
		person:     person{name: "John", age: 30},
		employeeID: "E123",
		salary:     60000.0,
	}

	fmt.Println("Name:", emp.name)
	fmt.Println("Age:", emp.age)
	fmt.Println("Employee ID:", emp.employeeID)
	fmt.Println("Salary:", emp.salary)
	emp.greet()
}
