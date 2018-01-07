package main

import (
	"net/http"
	"fmt"
)

var urls = []string{
	"http://tccbest.com/",
	"http://baidu.com/",
	"http://sina.com/",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("ERROR: ", url, err)
		}

		fmt.Println(url, ": ", resp.Status)
	}
}
