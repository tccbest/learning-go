package main

import (
	"net/http"
	"fmt"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello!")
}

func main() {
	var h hello
	http.ListenAndServe(":4000", h)
}
