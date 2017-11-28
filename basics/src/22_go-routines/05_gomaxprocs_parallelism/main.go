package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 4500; i++ {
		fmt.Println("Foo:", i*i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 4500; i++ {
		fmt.Println("Bar:", i*i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
