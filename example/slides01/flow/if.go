package flow

import "fmt"

func ifTrue() {
	// pay attention to the lack of brackets
	if true { // true-ism, that's it
		fmt.Println("I would be surprissed if you don't see a message")
	}
}

func ifComplex(isTrue bool) {
	if isTrue && sometimesTrue() { // math works as expected
		fmt.Println("Should be seen sometimes")
	}
}

func ifElse() {
	if false {
		fmt.Println("nothing...")
	} else {
		fmt.Println("Sometimes you are right even if you are wrong")
	}
}

func ifAndAssign() {
	if number := 42; number%2 == 0 {
		fmt.Println("See the message based on the calculated variable")
	} else {
		fmt.Printf("the scope of `number` variable is limited to the `if/else`")
	}
}

func sometimesTrue() bool {
	return true // that's it, no random generator for now
}
