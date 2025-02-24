package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/ws",hendleConnection)
	log.Fatal(http.ListenAndServe(":3030",nil))
}