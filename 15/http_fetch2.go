package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter thr url...")

	url, err := inputReader.ReadString('\n')
	url = strings.TrimSuffix(url, "\n")
	checkErr(err)
	res, err := http.Get(url)
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
