package structs

import "fmt"

// Empty's capital letter allows to export it out
type Empty struct{} // the form of minimal type declaration

// User is a typical example
type User struct {
	Name string // the same capitalization rule is used for fields
	Age  int    // and functions, constants, and variables if it maatters
}

func moreExamples() {
	realUser := User{ // multiline instantiation, can be moved into the 'constuctor' func if needed
		Name: "Nikita",
		Age:  27,
	}

	fmt.Println(realUser)
}

// example: do not pass age, use default value (0 for any numeric type)
func NewUser(name string) User {
	return User{Name: name} // inline form, `Age` is omitted
}

// it is possible to make `methods` for structs (hello `basic` OOP!)
func (u User) SayHello() string {
	return fmt.Sprintf("Hello, my name is %s, I am %d", u.Name, u.Age)
}
