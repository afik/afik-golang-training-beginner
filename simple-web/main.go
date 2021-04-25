package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world!")
}
