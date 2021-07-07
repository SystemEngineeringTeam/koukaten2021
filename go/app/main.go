package main

import (
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/koukaten2021/apifuncs"
)

func main() {
	http.HandleFunc("/people", apifuncs.Getpeople)
	http.ListenAndServe(":80", nil)
	log.Print("起動しました\n")
}
