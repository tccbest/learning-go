package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	who := "Awesome"
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}

	fmt.Println("Good Morning", who)
}
