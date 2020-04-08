package main

import "fmt"

type rect struct {
	width, height int
}

// these are the methods allowed on the above struct
func (r rect) area() int {
	return r.width * r.height
}

// go can handle the value and pointer conversion
// use pointers if u want to mutate the struct
// else simple passing would do like the above case
func (r *rect) perim() int {
	return (r.width + r.height)*2
}

func main() {
	r := rect{ width: 10, height: 20 }
	fmt.Println("width ", r.area())
	fmt.Println("height ", r.perim())
}