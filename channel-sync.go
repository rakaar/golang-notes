package main 

import ("fmt"
	"time"
)

func learn_block(done chan<- bool) { // chan<- is used to specify the direction of the data: either to or from 
	// <-chan means only data can be read from the channel
	fmt.Println("start")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	sample_chan := make(chan bool, 1)
	go learn_block(sample_chan)
	
	// if the below line weren't there the main program would exit
	// but here the below line has blocked until, the current routine(main) will exit only after receving the data to the channel
	<-sample_chan
}