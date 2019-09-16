package concurrency

import (
	"fmt"
	"time"
)

// struct{} is the simplest piece of data you can send with channel
// it is used just a confirmation, no real data is obviosly passed there
func taskPerformer(syncDone chan struct{}) {
	fmt.Println("start work")
	time.Sleep(time.Second)
	fmt.Println("finish work")

	// pass an empty struct just to pass the reading from the channel
	syncDone <- struct{}{}
}

func exampleWithSync() {
	// init the channel
	done := make(chan struct{})

	// start the work, pass the sync channel
	go taskPerformer(done)

	// perform some other work in main goroutine
	fmt.Println("main goroutine work")
	time.Sleep(200 * time.Millisecond)

	// wait until the async job is done
	<-done // perform the reading from the channel
	fmt.Println("Job's done!")
}
