package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//	type Guestbook struct {
//		SignatureCount int
//		Signatures []string
//	}
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	//message := []byte("Hi^ men!")
	signatures := getStrings("signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, nil)
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
	http.HandleFunc("/", viewHandler)
	_, err := os.Stdout.Write([]byte("server created in http://localhost:4000/"))
	check(err)

	err = http.ListenAndServe("localhost:4000", nil)
	log.Fatal(err)
}
