package main

import ("fmt"
	 "errors"
)

func div(arg int) (int, error) {
	// handle errors
	if arg == 0 {
		return -1, errors.New("Cant divide with zero")
	} else {
		return 1, nil
	}
}

func main() {
	fmt.Println(div(0))
}