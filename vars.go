package main

import "fmt"

func main() {
	var a = 1;
	fmt.Println(a);

	// the below is not allowed as we already used 'a' as a int
	// a = "not allowrd"
	// Println(a)

	// in go u need no mention the type as shown above, though u can likr
	var test string = "hello"
	fmt.Println(test)

	// the simplest for declaring as well as initialising is using ":=" this notaion
	simple := "can be anything"
	fmt.Println(simple);

	// if not initalized corressponds to its zero values 
	var x int
	var y string
	fmt.Println(x) // 0
	fmt.Println("this is empty ", y) // ""
}