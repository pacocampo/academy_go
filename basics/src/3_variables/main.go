package main

import "fmt"

type lat float32
type lon float32

var name string
var age int
var family bool
var latitude lat = 9.1355
var longitude lon = 90.34356

func main() {
	lastname := "ocampo"
	fmt.Printf("%v\n", name)
	fmt.Printf("%v\n", age)
	fmt.Printf("%v\n", family)
	fmt.Printf("%v\n", lastname)
	fmt.Printf("%T\n", latitude)
	fmt.Printf("%T\n", longitude)
}