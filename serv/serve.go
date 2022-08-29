package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

type info struct {
	db *sql.DB
}

//отправляем в базу данных значение из формы
func (dbStruct *info) createHandlerNew(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")

	_, err := dbStruct.db.Exec(`INSERT INTO signatures (phrase) VALUES($1)`, signature)

	if err != nil {
		log.Println("error of insert:", signature, err)
	}
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}
func addSignatureHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("ui/form.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	dsn := "postgresql://corkiudy:test@127.0.0.1:5432/forProjectGo?sslmode=disable"
	db, err := openDB(dsn) // инициализируем пул подключений к базе
	defer db.Close()

	dbStruct := &info{
		db: db,
	}
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", addSignatureHandler)
	http.HandleFunc("/guestbook/create", dbStruct.createHandlerNew)
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
