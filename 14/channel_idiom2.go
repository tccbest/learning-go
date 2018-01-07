package main

import (
	"fmt"
	"time"
)

func pupm() chan int  {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++  {
			ch <- i
		}
	}()

	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func main() {
	suck(pupm())
	time.Sleep(1e9)
}
