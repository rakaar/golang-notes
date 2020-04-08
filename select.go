package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	} ()

	go func() {
		time.Sleep(2*time.Second)
		c2 <- "two"
	}()

	// it is used to await for both of them
	//A select blocks until one of its cases can run, then it executes that case.
	// It chooses one at random if multiple are ready.
	for i := 1; i <= 2; i++{
		select {
		case msg1 := <- c1:
			fmt.Println("suces ", msg1)
		case msg2 := <- c2:
			fmt.Println("Suces ",msg2)
		}
	}
	
}