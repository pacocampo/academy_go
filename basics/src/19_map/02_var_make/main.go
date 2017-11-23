package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	myGreeting["Miguel"] = "Días!"

	fmt.Println(myGreeting)
}
