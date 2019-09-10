package interfaces

import "fmt"

// DoSomething is just like an interface from other languages,
// it promotes some type of behavior
type Worker interface {
	Work() string
}

// RealWorker struct and it's Work method __implicitly__ suffice the interface
type RealWorker struct {
	Name string
}

// Work function declaration is the only thing that is required
// in order to implement the interface.
func (w RealWorker) Work() string {
	return "allright! I will work! Job's done!"
}

// just a small constructor function
func newRealWorker() RealWorker {
	return RealWorker{} // workers have no name!
}

func example() {
	// we instantiate __concrete__ type
	worker := newRealWorker()

	// and pass it as interface type
	forceToWork(worker)
}

// here we expect anything that implements `Worker` interface (or just has `Work` function)
func forceToWork(worker Worker) {
	fmt.Println("it works!")
	fmt.Printf("it ways something: %s!", worker.Work())
}
