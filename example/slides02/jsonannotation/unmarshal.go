package jsonannotation

import (
	"encoding/json"
	"fmt"
)

func exampleFromJSON() {
	// prepare the string to parse (from the previous example)
	userJSON := []byte(`{"name":"Alex","Age":24,"professtion":"singer"}`)

	var user User                   // prepare the variable where to put the parsed data
	json.Unmarshal(userJSON, &user) // unmarshal: pay attention that the pointer is passed!
	// the pointer is required in order to update the internal fields in the outer scope

	// print the result with fields
	fmt.Printf("the parsed user: %+v", user)
}
