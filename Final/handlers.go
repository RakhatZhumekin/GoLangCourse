package main

import "net/http"

func Main(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Hello"))
}