package errorhandling

import (
	"fmt"
	"time"
)

// CustomError provides more details on possible error and
// as the error is just an interface, we can implement our own errors, just like exceptions!
type CustomError struct {
	HappenedAt  time.Time
	Cause       string
	Responsible string
}

// NewCustomError returns an actual error
func NewCustomError(cause, responsible string) error {
	return CustomError{
		Cause:       cause,
		Responsible: responsible,
		HappenedAt:  time.Now(),
	}
}

// just a simple Error() implementation
func (ce CustomError) Error() string {
	return fmt.Sprintf("%s: %s happened because of %s",
		ce.HappenedAt, ce.Cause, ce.Responsible)
}

// the function that uses the error
func dangerousCustomFunction() error {
	return NewCustomError("test", "Nikita")
}

func exampleCustomError() {
	// the error handling is the same
	err := dangerousCustomFunction()
	if err != nil {
		fmt.Printf("oh no, the error! Details: %s", err)
	}
}
