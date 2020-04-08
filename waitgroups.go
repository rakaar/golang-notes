package main

import (
	"fmt"
	"sync"
	"time"	
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker started  ", id)

	time.Sleep(time.Second)
	fmt.Println("Worker ended ", id)
	
}

func main() {
	// this waitgroup  will wait for all go routines to finish
	var wg sync.WaitGroup

	// launching go routines and launching  the counter
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// blocks until waitgroup goes back to zero
	wg.Wait()
}