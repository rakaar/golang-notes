package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("res status ", res.Status)
}