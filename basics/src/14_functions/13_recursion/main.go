package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	fmt.Println(x)
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}
