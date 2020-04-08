package main

import "fmt"

func main() {
	// looping
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// if else
	x := 3;
	if (x == 1) {
		fmt.Println("In if")
	} else if (x == 2) { // SPECIAL  NOTE THAT else if should be in the same line as in ending if brace 
		fmt.Println("in else if")
	} else {
		fmt.Println("in else")
	}

	// arrays
	var a[3] int 
	a[1] = 100
	fmt.Println(len(a))
	fmt.Println(a[0]) // 0

	// slices : arrays + more features
	slice1 := make([]int, 3)
	slice1[0] = 3
	fmt.Println(slice1[0]) // 3
	// append	
	slice2 := append(slice1, 4) 
	fmt.Println(slice2[3]) // 4
	// copy
	slice3 := make([]int, 4)
	copy(slice3, slice2)
	fmt.Println(slice3[3]) // 4
	fmt.Println(slice3)
	
	// maps
	map1 := make(map[string]int)
	map2 := map[string]int{"hello": 1, "hi": 2}
	
	map1["hello"] = 1
	map1["hii"] = 2
	
	fmt.Println(map2) // map[hello:1 hi:2]

	delete(map1, "hello")
	fmt.Println(map1) // map[hii:2]

	// range: iterate over arrays and maps
	nums := [3]int{1,2,3}
	for ind, value := range nums {
		fmt.Println("index ",ind, "value", value)
	}

	for key, value := range map2 {
		fmt.Println("key ", key, "value ", value)
	}
	
} 