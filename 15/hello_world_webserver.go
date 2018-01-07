package main

import (
	"net/http"
	"fmt"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("in")
	fmt.Fprintf(w, "helle," + req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("li", err.Error())
	}
}
