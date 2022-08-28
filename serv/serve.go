package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hi^ men!")
	_, err := writer.Write(message)
	check(err)
}

func main() {
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe("localhost:4000", nil)
	log.Fatal(err)
}
