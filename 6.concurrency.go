package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type response struct {
	Website   string
	Status    string
	result    string
	timeTaken int64
}

func request(website string, channel chan response, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	res, err := http.Get(website)
	requiredTime := time.Since(start).Milliseconds()
	if err != nil {
		channel <- response{Website: website, Status: res.Status, result: err.Error(), timeTaken: requiredTime}
		return
	}
	body, err := io.ReadAll(res.Body)
	// buffer := make([]byte, 1024)
	// n, err := res.Body.Read(buffer)
	if err != nil {
		channel <- response{Website: website, Status: res.Status, result: err.Error(), timeTaken: requiredTime}
		return
	}
	defer res.Body.Close()
	channel <- response{Website: website, Status: res.Status, result: string(body), timeTaken: requiredTime}
}

// func main() {
	// **Concurrency**
	// Go has built-in support for concurrent programming using goroutines, wait groups and channels.
	// Channels are used to communicate between goroutines and wait groups are used to wait for multiple goroutines to finish.
	// Goroutines are lightweight threads managed by the Go runtime and can access as many threads as the system can handle.
	// They are created using the `go` keyword followed by a function call.

	// Buffered channels allow sending and receiving without blocking until the buffer is full.
	// Unbuffered channels block until both sender and receiver are ready.
	// Buffered channels can be created by passing a capacity to the `make` function.
	channel := make(chan response) // Create a channel of type int, unbuffered by default and channel is a pointer
	var wg sync.WaitGroup          // Create a wait group to wait for goroutines to finish

	websiteList := []string{"https://pkg.go.dev", "https://google.com", "https://github.com/ISanviI", "https://stackoverflow.com", "https://reddit.com"}

	totalTime := int64(0)
	goRuntime := time.Now()
	for i := 0; i < len(websiteList); i++ {
		wg.Add(1)
		go request(websiteList[i], channel, &wg)
		// One could also receive message from channel here if it is unbuffered.
	}
	goEnd := time.Since(goRuntime).Milliseconds()
	fmt.Printf("Goroutines started at: %s and ended in %d ms.\n", goRuntime.Format(time.RFC3339), goEnd)

loop:
	for {
		// `select` is used to wait on multiple channel operations. It will block until one of the cases can proceed.
		// It is similar to `switch` but for channels.
		select {
		case msg := <-channel: // receive from data channel
			totalTime += msg.timeTaken
			fmt.Printf("Website: %s, Status: %s, Time Taken: %d ms\n", msg.Website, msg.Status, msg.timeTaken)

		case <-time.After(5 * time.Second): // wait for 5 sec inactivity
			fmt.Println("no more messages, exiting...")
			break loop

			// Including a default here to avoid blocking if no messages are available, however in this case if we also add default, the inactivity case would never be executed.
		}
	}
	fmt.Printf("\nTotal time taken for all websites with go runtime: %d ms\n", totalTime)

	// If the channel had been closed by sender after sending all messages:
	// msg, ok := <- channel {
	// 	if !ok {
	// 		fmt.Println("Channel is closed, no more messages will be received.")
	// 		break
	// 	}
	// 	totalTime += msg.timeTaken
	// }
	// OR
	// for msg := range channel {
	// 	totalTime += msg.timeTaken
	// }
	// Sending on a closed channel causes panic!!
	// If the channel is unbuffered, the sender should always close the channel after sending all messages.
	// If the channel is buffered, it can be closed by the sender after sending all messages, but it is not necessary.

	allWebsTime := int64(0)
	for i := 0; i < len(websiteList); i++ {
		start := time.Now()
		_, err := http.Get(websiteList[i])
		if err != nil {
			fmt.Printf("Error fetching %s: %s\n", websiteList[i], err.Error())
			continue
		}
		allWebsTime += time.Since(start).Milliseconds()
	}
	fmt.Printf("\nTotal time taken for all websites without go runtime: %d ms\n", allWebsTime)

	defer close(channel) // The sender should always close the channel
	wg.Wait()            // Wait for all goroutines to finish
}

// Tickers in GO that return channels
// 1. time.Tick() - Returns a channel that sends the current time at regular intervals.
// 2. time.After() - Returns a channel that sends the current time after a specified duration.
// 3. time.Sleep() - Pauses the current goroutine for a specified duration.

// Reading from a NIL channel - Blocks the receiver forever
// Writing to a NIL channel - Causes deadlock as it blocks the sender forever
// Reading from a closed channel - Returns the zero value of the channel type and a boolean false indicating that the channel is closed.
// Writing to a closed channel - Causes panic
