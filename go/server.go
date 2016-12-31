package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

/**********************SQLITE TESTING ********************************************************************/

type TestItem struct {
	Id    string
	Name  string
	Phone string
}

//Open the db object
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS items(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT,
		Phone TEXT,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func StoreItems(db *sql.DB, items []TestItem) {
	sql_additem := `
	INSERT OR REPLACE INTO items(
		Id,
		Name,
		Phone,
		InsertedDatetime
	) values(?, ?, ?, CURRENT_TIMESTAMP)
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(item.Id, item.Name, item.Phone)
		if err2 != nil {
			panic(err2)
		}
	}
}

func ReadItem(db *sql.DB) []TestItem {
	sql_readall := `
	SELECT Id, Name, Phone FROM items
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []TestItem
	for rows.Next() {
		item := TestItem{}
		err2 := rows.Scan(&item.Id, &item.Name, &item.Phone)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

func Database() {
	db := InitDB("./sqlite.db")
	CreateTable(db)
	var items = []TestItem{{"id", "Name", "8675309"}}
	StoreItems(db, items)
	fmt.Print(ReadItem(db))
}

/**********************SQLITE TESTING ********************************************************************/

//Make into a struct
func DatabaseHandler(w http.ResponseWriter, r *http.Request) {

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}

func LightHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println("Light Handler hit " + id)
}

func PumpHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println("Pump Handler hit " + id)
}

func Serve() {
	var host = "localhost"
	var port = "8080"

	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/light/{id}", LightHandler)
	// r.HandleFunc("/pump/{id}", PumpHandler)
	http.Handle("/dist/", http.StripPrefix("/dist", http.FileServer(http.Dir("../dist"))))
	// http.Handle("/", r)
	log.Printf("goLang server listening on %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
