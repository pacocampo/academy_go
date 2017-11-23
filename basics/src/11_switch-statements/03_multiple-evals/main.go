package main

import "fmt"

func main() {
	switch "Susan" {
	case "Tim", "Jenny", "Susan":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	default:
		fmt.Println("nothing to do")
	}
}
