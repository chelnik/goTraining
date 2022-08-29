package main

import (
	"bufio"
	"fmt"
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
		log.Fatal("Error:", err)
	}
}

func (dbStruct *info) viewHandler(writer http.ResponseWriter, request *http.Request) {
	//signatures := getStrings("signatures.txt")
	signatures := dbStruct.getStringsNew()
	html, err := template.ParseFiles("ui/index.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

func (dbStruct *info) getStringsNew() []string {
	// Query для запросов к базе
	rows, err := dbStruct.db.Query(`SELECT phrase FROM SIGNATURES`)
	check(err)
	//defer rows.Close() когда-нибудь ее нужно будет закрыть
	var arrString []string
	for i := 0; rows.Next(); i++ {
		var s string
		err = rows.Scan(&s)
		arrString = append(arrString, s)
		check(err)
	}
	fmt.Println(arrString)
	return arrString
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
