package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// Esto origina un deadlock.
// porqué?
// cómo arreglarlo?
