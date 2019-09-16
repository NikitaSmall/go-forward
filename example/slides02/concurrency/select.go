package concurrency

import (
	"fmt"
	"time"
)

// a job to start when the message arrives
func processA(data string) {
	fmt.Printf("processing A with %s\n", data)
	time.Sleep(time.Second)
}

// another job to perform
func processB(data string) {
	fmt.Printf("processing B with %s\n", data)
}

func exampleWithSelect() {
	// data channels
	channelA := make(chan string)
	channelB := make(chan string)

	go sendDataToChannel(channelA, channelB)

	for messagesRevieved := 0; messagesRevieved < 10; messagesRevieved++ {
		// wait and perform any of the task whichever comes first (blocking execution)
		select {
		case data := <-channelA:
			processA(data)
		case data := <-channelB:
			processB(data)
			// in order to make the select non-blocking and passable over time
			// default:
			// 	fmt.Println("no activity in both channels")
			//  time.Sleep(time.Second)
		}
	}
}

func sendDataToChannel(chanA, chanB chan string) {
	num := 0
	for {
		chanA <- fmt.Sprintf("message a number %d", num)
		// even if we have the second message immideately, we still wait for the execution
		// in select, we perform the task in one thread
		chanB <- fmt.Sprintf("message b number %d", num)
		time.Sleep(time.Second)

		num++
	}
}
