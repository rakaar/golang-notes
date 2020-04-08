package main

import "os"

func main() {
	panic("Issue")

	_, err := os.Create("/ffgg")
	if err != nil {
		panic(err)
	}
}