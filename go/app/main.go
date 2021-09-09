package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"set1.ie.aitech.ac.jp/koukaten2021/apifuncs"
	"set1.ie.aitech.ac.jp/koukaten2021/camera"
)

func main() {
	log.Print("起動しています...\n")
	http.HandleFunc("/people", apifuncs.Getpeople)

	go camera.CameraTimer()
	http.ListenAndServe(":80", nil)
}
