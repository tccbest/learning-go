package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	a := "我是Awesome哈哈"

	fmt.Println(len(a), utf8.RuneCountInString(a))

	for _, val := range a {
		fmt.Printf("%s\n", string(val))
	}

	c := []byte(a)
	c[0] = 'a'
	fmt.Println(string(c))

	a = strings.Replace(a, "我", "你", 1)
	fmt.Println(a)
}
