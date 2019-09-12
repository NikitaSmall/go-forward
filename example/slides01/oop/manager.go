package oop

import "fmt"

// Manager struct embeds Employee and thus has access to Employee's methods and fields
type Manager struct {
	Employee      // embedding a nameless Employee here
	managedSector string
}

// can't pass non-Employee
func expectPureEmployee(emp Employee) {
	fmt.Printf("internals: %s", emp.format())
}

// implement an interface
func (m Manager) roleName() string {
	return fmt.Sprintf("manager on sector %s", m.managedSector)
}

func example2() {
	newManager := Manager{
		Employee: Employee{
			id:     1,
			name:   "Steeve",
			salary: 1000,
		},
		managedSector: "IT",
	}

	newManager.Greeting() // it speaks just as Employee
	// but is actually the same as
	newManager.Employee.Greeting()

	// won't work:
	// expectPureEmployee(newManager)
}
