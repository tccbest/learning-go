package main

import "fmt"

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in
			if i%prime != 0 {
				out <- i
			}
		}
	}()

	return out
}

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()

	return out
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
