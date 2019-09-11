package errorhandling

import (
	"errors"
	"fmt"
)

func example() error {
	res, err := typicalErrorReturn()
	if err != nil {
		fmt.Printf("do something! There is an error: %s", err)

		// usually there is an exit/return or actual error handling
		return err
	}

	// proceed if everything is fine
	fmt.Println(res)

	if err := neverAnError(); err != nil {
		//handle an error, possibly add details
		return err
	}

	return nil
}

func neverAnError() error {
	// assume the work is done and no error happens
	return nil
}

func typicalErrorReturn() (string, error) {
	return "function returns a value/result", errors.New("possible error")
}
