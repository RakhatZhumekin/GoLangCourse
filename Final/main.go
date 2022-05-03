package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Main)

	log.Fatal(http.ListenAndServe(":11037", nil))
}