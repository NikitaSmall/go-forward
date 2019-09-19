package jsonannotation

import (
	"encoding/json" // the standard library that allows json related transformation
	"fmt"
)

// run this https://play.golang.org/p/Eyc3lmJEpiG
func exampleToJSON() {
	user := NewUser("Alex", "singer", "male", 24)

	// just a json package Marshal function call, nothing else
	// it returns a byte slice (array) and error
	jsonResult, err := json.Marshal(user)
	if err != nil {
		// if the error is not nil, then something caused an error
		// this is the simpliest way to handle the error (more on exceptions later)
		fmt.Printf("something bad happened! the error: %e \n", err)
	}

	// cast byte slice to a string and print it (check the json)
	fmt.Println(string(jsonResult))
}
