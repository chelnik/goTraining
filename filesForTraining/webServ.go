package main

import (
	"fmt"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "This is main page!")
}
func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	fmt.Printf("create web-server on http://%s", "localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

