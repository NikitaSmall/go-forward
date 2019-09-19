package jsonannotation

// annotation is a way to specify some additional information about the object
// this is should be done only during the development process but there is a way
// to set the annotation in runtime.
// the information can be used for marshaling (json, xml), DB interactions
// whole bunch of libraries use the annotations as a way to set some config
// in a really neat and simple way.

// User is a great annotation example:
// basically, this is all you need in order to correctly marshal the struct
type User struct {
	// you can only marshal exported fields (the capitalized)
	Name string `json:"name"` // set the name to marshal the Name field from/to JSON
	Age  int    // the default JSON name is the same as the field (pay attention to the capital letter)

	SecretCode string `json:"-"`                     // do not export field
	Profession string `json:"professtion,omitempty"` // omit the field in JSON if it's empty

	gender string // no marshal, as the field is private (isn't exported)
}
