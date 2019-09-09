package assignment

import "log"

func assign() {
	var variable string = "this is a variable"

	otherVar := "you can ommit type if it's clear to the compiler"
	typesOfAssignment := "do not mix `var ... type` and `:=`, they don't work together"
	var (
		massAssign string = "you can assign a number of variables if you wish"
		third      string = "as this still is a `var ... type` assignment"
	)

	const singleConst string = "this will not change"
	const (
		massConst string = "is possbile"
	)

	caution := "you have to use the variable once it's declared!!!"
	log.Print(variable, otherVar, typesOfAssignment,
		massAssign, massConst, third, caution)
}
