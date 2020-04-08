package main

import ("fmt"
		"time"
)

func do(word string) {
	fmt.Println(word)
}



func main() {
	go do("routines")

	time.Sleep(time.Second) // this sleeps the main routine so as to give time for the above routine to execute

	do("main")

	msg := make(chan string)
	go func() { msg <- "ping" }() // writing to a channel

	word := <- msg 	// reading from a channel
	// this go routine waits until data is read to the channel

	fmt.Println("word is ", word)

	// making a buffered channel
	buf := make(chan int, 2)
	buf <- 10
	buf <- 20

	fmt.Println("buf 1 ", <-buf)
	fmt.Println("buf 2 ", <-buf)


	
}