package oop

import "fmt"

// Employee is a struct that holds data
type Employee struct {
	id     int // first lowercase letter means that the field is no exported
	name   string
	salary float32
}

// method that modifies the internal field
// pay attention on the pointer
func (e *Employee) raiseSalary(bonus float32) {
	e.salary += bonus
}

func (e Employee) format() string {
	return fmt.Sprintf("Id: %d, Name: %s, Salary: %.2f", e.id, e.name, e.salary)
}

// Greeting is a public method that has any instance of Employee struct
func (e Employee) Greeting() { // public because of uppercased first letter
	fmt.Printf("Hello: %s", e.name)
}

// implement an interface
func (e Employee) roleName() string {
	return "Generic employee"
}

func example1() {
	newEmployee := Employee{
		id:     42,
		name:   "Bob",
		salary: 300,
	}

	newEmployee.Greeting()
	newEmployee.raiseSalary(250)
}
