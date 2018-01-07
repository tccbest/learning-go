// test_oddeven.go
package main

import (
	"fmt"
	"learning-go/13o/13/even/even"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
