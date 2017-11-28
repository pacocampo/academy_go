package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		fmt.Println("primer done")
		<-done
		fmt.Println("segundo done")
		close(c)
		fmt.Println("voy a cerrar")
	}()

	for n := range c {
		fmt.Println(n)
	}
}
