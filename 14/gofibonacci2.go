package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i <= n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {
	start := time.Now()

	c := make(chan int, 25)

	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
