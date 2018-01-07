package main

import (
	"net/http"
	"fmt"
	"strings"
)

func helloHandler(w http.ResponseWriter, req *http.Request)  {
	url := req.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "hello %s", url)
}

func shouthelloHandler(w http.ResponseWriter, req *http.Request)  {
	url := req.URL.Path[len("/shouthello/"):]
	fmt.Fprintf(w, "hello %s", strings.ToUpper(url))
}

func main()  {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	http.ListenAndServe(":8080", nil)
}
