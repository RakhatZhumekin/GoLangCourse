package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var database = NewRedis()

func main() {
	fmt.Println("Starting the REST server on port 11037")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/put", put).Methods("PUT")
	router.HandleFunc("/get/{key}", get).Methods("GET")

	if error := http.ListenAndServe(":11037", router); error != nil {
		fmt.Println(error.Error())
	}
}

func put(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode("Couldn't read the body")
		return
	}

	requestText := string(requestBody[:])

	request := strings.Split(requestText, " ")

	if len(request) != 3 || strings.ToLower(request[0]) != strings.ToLower("set") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode("Incorrect syntax. Example: set key value")
		return
	}

	database.Put(request[1], request[2])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode("Successfully added")
}

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]

	value := database.Get(key)

	if value == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode("No such entry with the given key")
		return
	}

	pair := make(map[string]string)
	pair[key] = value

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pair)
}