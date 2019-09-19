package errorhandling

import "fmt"

func examplePanic() {
	// sometimes the error can't be handles, just like an exception
	// then we should explicitly tell the program to ... PANIC! (this is usually bad)
	panic("the reason of the panic or responsible value")

	// caution: this will break and stop the program unless handled
}

// panic shouldn't be use to replicate try/catch (use if-error instead)
// but sometimes you should be able to fail explicitly
// try not to panic at first.
func exampleRecover() {
	n := panicFunction()
	fmt.Println("main received", n)
}

func panicFunction() (m int) {
	defer func() {
		// recover function forces panic to stop
		// thus, we still will be able to proceed with the execution
		if err := recover(); err != nil {
			fmt.Println(err) // this block can be reached only if the panic happened
			m = 2
		}
	}()
	m = 1
	panic("panicFunction: panic example")
	// m = 3 nothing reaches this point
	// return m
}
