package channel

import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	ch <- "hello there!" // send a message from the channel
}

func reciever(ch chan string) {
	// read from channel
	// such read will block is no data is available in the channel
	message := <-ch
	fmt.Printf("get the message: %s\n", message) // use the result
}

func example() {
	// declare and initialize channel
	messageChannel := make(chan string) // make should be done

	// simple data transfer, should be from or to another goroutine
	go sender(messageChannel)
	reciever(messageChannel)

	// channels are very useful when used with goroutines
	go func() {
		for { // send messages every second
			sender(messageChannel)
			time.Sleep(1 * time.Second)
		}
	}()

	// wait for and recieve five messages, then exit (loop ends)
	for i := 0; i < 5; i++ {
		reciever(messageChannel)
	}
}
