package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func update1(id int, counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	// Counter could be a closure too.
	defer wg.Done()

	for i := 0; i < 3; i++ {
		// lock before updating shared var
		mu.Lock()
		*counter++
		fmt.Printf("Worker %d incremented counter to %d\n", id, *counter)
		mu.Unlock()

		// simulate work
		time.Sleep(200 * time.Millisecond)
	}
}

func update2(id int, squareSum *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	*squareSum += id * id
	mu.Unlock()
	fmt.Printf("Worker %d is done\n", id)
	time.Sleep(100 * time.Millisecond)
}

func reader(id int, c1, c2 *int, mu *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mu.RLock() // shared lock
		fmt.Printf("(Reader %d) sees counters: c1=%d, c2=%d\n", id, *c1, *c2)
		mu.RUnlock()

		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
}

// func main() {
	var (
		counter1 int
		mu1      sync.Mutex
		wg       sync.WaitGroup
		mu2      sync.Mutex
		counter2 int
		rmu      sync.RWMutex
		rwg      sync.WaitGroup
	)

	// spawn multiple goroutines calling the same function
	for w := 1; w <= 3; w++ {
		wg.Add(2)
		go update1(w, &counter1, &mu1, &wg)
		go update2(w, &counter2, &mu2, &wg)
	}

	for r := 1; r <= 3; r++ {
		wg.Add(1)
		go reader(r, &counter1, &counter2, &rmu, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished.")
	// Final values of counters
	for r := 1; r <= 3; r++ {
		rwg.Add(1)
		go reader(r, &counter1, &counter2, &rmu, &rwg)
	}
	rwg.Wait()

	fmt.Println("Final Counter:", counter1)
	fmt.Println("Final Square Sum:", counter2)
}
