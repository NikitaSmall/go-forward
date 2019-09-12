package goroutine

import (
	"fmt"
	"time"
)

func call(message string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Printf(message)
	}
}

func example() {
	call("sync call") // this will wait three seconds and write a message after each second

	go call("async call") // goroutine will run the same time in total
	// but the rest of the execution will not be blocked

	// goroutines could be used with clousures
	externalMessage := "message"
	go func() {
		time.Sleep(time.Second)
		fmt.Printf(externalMessage)
	}() // define and call the closure

	time.Sleep(3 * time.Second) // let's wait for the result
}
