package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	res, err := http.Get("http://tccbest.com")
	checkErr(err)

	data, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	fmt.Printf("Got: %q", string(data))
}

func checkErr(err error)  {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
