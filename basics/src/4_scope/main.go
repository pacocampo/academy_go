package main

import (
	"4_scope/demo"
	"fmt"
)

func main() {
	x := 15
	fmt.Println(x)
	demo.MyName = "otro nombre"
	fmt.Println(demo.FullName())
}

