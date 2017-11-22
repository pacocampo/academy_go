package main

import "fmt"

const e string= "hello"

const (
	Pi = 3.1416
	Language = "Go"
)

const (
	a = iota * 20
	b
	c
	d
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
}