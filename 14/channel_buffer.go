package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan int)
	go func() {
		time.Sleep(1 * 1e9)
		x := <-c
		fmt.Println("RECEIVED", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)

	//time.Sleep(1e9)
}
