package main

import (
	"fmt"
	"log"
	"net/http"
)

func requestHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hi there, I love %s!", req.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", requestHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}