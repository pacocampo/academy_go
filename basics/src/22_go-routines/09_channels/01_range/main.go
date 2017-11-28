package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(time.Second)
		}

		// close(c)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			// time.Sleep(time.Second)
		}

		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
