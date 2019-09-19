package errorhandling

// Golang doesn't have exception mechanism in a typical way
// so if the function behaves dangerously, the possible error should be handled
// when the error marker is returned

// it's pretty common to return two values: actual result and error
// error is just another value. Usually, it's a string message wrapped in the interface
import (
	"errors"
	"fmt"
)

// this function returns an error interface which means it returns nil if everything is fine
func dangerousFunction() error {
	// the simplest way to create an error: use `errors` package
	return errors.New("Just a simple error description")
}

func exampleProcessing() {
	err := dangerousFunction()
	if err != nil { // it's the most common error handling pattern
		// we know this will always fire
		fmt.Println("oh no, an error!")
	}

	// some other work if everything is fine
}
