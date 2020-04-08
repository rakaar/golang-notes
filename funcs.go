package main

import "fmt"

func add(a, b int) int{
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

// multiple return values
func both(a, b int) (int,int) {
	return a + b,a - b
} 

// variable arguments
func varargs(nums ...int) {
    fmt.Println(nums, " ")
  
}

// closures , function returning another annonymus funcs

func intSeq() func() int {
	i := 0
	return func () int {
		i++
		return i
	}
}

// pointer
func incr(ptr *int) {
	*ptr = *ptr + 1
}

func main() {
	result := add(1,2)
	fmt.Println(result)

	sum, diff := both(2,1)
	fmt.Println(sum, diff) // 3,1

	varargs(1,2)

	// using something similar to spread
	nos := []int{1,2,3,4,5,6}
	varargs(nos...)

	// assign the return result of intSeq to a variable nextInt
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2

	nextInt2 := intSeq()
	fmt.Println(nextInt2()) // 1

	test_num := 111
	incr(&test_num)
	fmt.Println(test_num)

}