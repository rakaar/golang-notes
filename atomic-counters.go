package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func main() {
	var ops uint64

	var wg sync.WaitGroup 

	for i := 1; i <= 30; i++ {
		wg.Add(1)

		go func () {
			for c := 1; c <= 100; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}


	wg.Wait()

	fmt.Println("ops is ", ops) // 30 x 100
}