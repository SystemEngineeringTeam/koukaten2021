package main

import (
	"log"
	"net/http"
)

func hoge(http.ResponseWriter, *http.Request) {
	log.Print("Accessed /")
}

func main() {
	log.Print("hello world\n")
	http.HandleFunc("/", hoge)

	http.ListenAndServe(":80", nil)
}
