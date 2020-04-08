package main

import (
	"fmt"
	"time"
)
func worker(id int, jobs <- chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started ", j)
		time.Sleep(time.Second) // to replicate an expensive job
		fmt.Println("worker ", id, "done ", j)
		results <- j*2
	}
}
func main() {
	// channels req for workers
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs and close the channel
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// wait for all the results
	for a := 1; a <= 5; a++ {
		<- results
	}
}