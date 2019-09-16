package concurrency

import (
	"fmt"
	"time"
)

func worker(id int, taskData <-chan int, resChan chan<- int) {
	// try to process the data from the channel until it closed
	for data := range taskData {
		fmt.Println("worker", id, "started  job", data)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", data)
		resChan <- data * 2 // return the result via the channel
	}
}

func exampleWithWorker() {
	// buffered channel will not block on writing to it if the channel has less than 100 objects in it
	taskData := make(chan int, 100)
	results := make(chan int, 100)

	for wID := 1; wID <= 3; wID++ { // run three workers
		go worker(wID, taskData, results)
	}

	dataLen := 10
	for i := 0; i < dataLen; i++ {
		taskData <- i
	}

	// channel closing blocks any writing to it
	close(taskData)

	for i := 0; i < dataLen; i++ {
		fmt.Printf("The result: %d \n", <-results)
	}

	fmt.Println("Job's done!")
}
