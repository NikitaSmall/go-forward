package simplicity

import "fmt"

// Counter allows counting the object
type Counter struct {
	count int    // default value for numeric type is zero
	desc  string // default value for string is empty string
}

func (c *Counter) increment() {
	c.count++ // just modify (increment)
}

func (c Counter) String() string { // implement a Stringer interface, allow graceful checks
	return fmt.Sprintf("%s: %d", c.desc, c.count)
}

func exampleDefault() {
	// pay attention, the count isn't initialized here (it's default)
	counter := Counter{
		desc: "example counter",
	}

	fmt.Println(counter) // check gracefully because of String()

	counter.increment()
	fmt.Println(counter)

	counter.increment()
	counter.increment()
	fmt.Println(counter)
}
