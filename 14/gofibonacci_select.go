package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	start := time.Now()

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i <= 25; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
