package funcs

import "fmt"

func simple() {
	fmt.Println("just do something")
}

func simpleVar(something string) {
	fmt.Printf("show something %s", something)
}

func returnIt(variable int) int {
	return variable + variable
}

func moreDifferentVars(str string, number int) {
	for i := 0; i < number; i++ {
		fmt.Println(str)
	}
}

func moreOfTheSame(name, book, desc string) string {
	return fmt.Sprintf("author: %s, book's name: %s, description: %s",
		name, book, desc)
}
