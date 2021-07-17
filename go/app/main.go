package main

import (
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/koukaten2021/apifuncs"
)

func main() {
	log.Print("起動しています...\n")
	http.HandleFunc("/people", apifuncs.Getpeople)
	http.ListenAndServe(":80", nil)
}
