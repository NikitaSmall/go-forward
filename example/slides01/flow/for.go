package flow

import "fmt"

func forEndless() {
	for { // this loop is endless and only `break` will stop it
		fmt.Println("The fun will never ends")
		// `break` or `continue` to exit or go next iteration for the loop
	}
}

func forAWhile(i int) {
	for i > 0 { // this is a way to set while loop in Go
		fmt.Printf("%d times left to loop", i) // by the way, hello string formatting!
		i--                                    // increment/decrement works as expected
	}
}

func forClassic(len int) {
	// classic, with the exception of no brackets. Go avoids them where possible
	for i := 0; i < len; i++ {
		fmt.Printf("Good old `for` with %d loops left", i)
	}
}
