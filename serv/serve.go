package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	_, err := writer.Write([]byte(signature))
	check(err)
}
func addSignatureHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("form.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	//message := []byte("Hi^ men!")
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}
func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", addSignatureHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	_, err := os.Stdout.Write([]byte("server created in http://localhost:4000/guestbook"))
	check(err)

	err = http.ListenAndServe("localhost:4000", nil)
	log.Fatal(err)
}
