package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

//type info struct {
//	db *sql.DB
//}

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
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

//func createHandlerNew(writer http.ResponseWriter, request *http.Request) {
//	signature := request.FormValue("signature")
//	dsn := flag.String("dsn", "postgresql://corkiudy:test@127.0.0.1:5432/forProjectGo?sslmode=disable", "Название postSQL источника данных")
//	db, err := openDB(*dsn) // инициализируем пул подключений к базе
//	if err != nil {
//		log.Println("Open db")
//	}
//	defer db.Close()
//	_, err = db.Exec(`INSERT INTO signatures (phrase) VALUES($1)`, signature)
//
//	if err != nil {
//		log.Println("go")
//	}
//	http.Redirect(writer, request, "/guestbook", http.StatusFound)
//}
func addSignatureHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("ui/form.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	//message := []byte("Hi^ men!")
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("ui/view.html")
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
	dsn := "postgresql://corkiudy:test@127.0.0.1:5432/forProjectGo?sslmode=disable"
	db, err := openDB(dsn) // инициализируем пул подключений к базе
	db.Close()

	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", addSignatureHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	_, err = os.Stdout.Write([]byte("server created in http://localhost:4000/guestbook"))
	check(err)
	err = http.ListenAndServe("localhost:4000", nil)
	log.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil { // проверка того что все настроено правильно
		return nil, err
	}
	return db, nil
}

//if err != nil {
//	log.Println("Open db")
//}
