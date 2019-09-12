package oop

import "fmt"

// Role is a generic simple interface that provides the work details
type Role interface {
	roleName() string
}

// just a check
func printRoleName(r Role) {
	fmt.Println("Role: ", r.roleName())
}

func test() {
	e := Employee{1, "Rafael", 1000.0}
	printRoleName(e) // e is implicitly converted to Role, kind of is-a relationship

	m := Manager{
		Employee: Employee{
			id:     1,
			name:   "Arnold",
			salary: 1000,
		},
		managedSector: "Sales",
	}

	printRoleName(m) // Manager implements Role too
}
