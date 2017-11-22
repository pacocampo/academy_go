package main

import "fmt"

// func main() {
// 	a := 10
// 	fmt.Println(a)
// 	fmt.Println(&a)

// 	var b *int = &a
// 	fmt.Println(b)
// 	fmt.Println(*b)

// 	*b = 15

// 	fmt.Println(a)
// 	fmt.Println(&a)
// }

func setToZero(x *int) {
	fmt.Println(&x)
	*x = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	setToZero(&x)
	fmt.Println(x)
}