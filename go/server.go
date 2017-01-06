package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

/**********************SQLITE TESTING ********************************************************************/
var db = InitDB("./sqlite.db")


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

func CreateTable() {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS events(
		Id TEXT NOT NULL PRIMARY KEY,
		Type TEXT,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func StoreEvents(events []Event) {
	sql_addevent := `
	INSERT OR REPLACE INTO events(
		Id,
		Type,
		InsertedDatetime
	) values(?, ?, CURRENT_TIMESTAMP)
	`

	stmt, err := db.Prepare(sql_addevent)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, event := range events {
		_, err2 := stmt.Exec(event.Id, event.Type)
		if err2 != nil {
			panic(err2)
		}
	}
}

func ReadEvents() []Event{
	sql_readall := `
	SELECT Id, Type FROM events
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Event
	for rows.Next() {
		event := Event{}
		err2 := rows.Scan(&event.Id, &event.Type)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, event)
	}
	return result
}

func initDatabase() {
	CreateTable()
	var events = []Event{{"id", "light"}}
	StoreEvents(events)
	fmt.Print(ReadEvents())
}

/**********************SQLITE TESTING ********************************************************************/


//Make into a struct
func DatabaseHandler(w http.ResponseWriter, request *http.Request) {

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

func CalendarHandler(w http.ResponseWriter, request *http.Request) {
	http.ServeFile(w, request, "../calendar.html")	
}

func PlantViewHandler(w http.ResponseWriter, request *http.Request) {
	http.ServeFile(w, request, "../plants.html")
}

func AddEventHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Add Event Handler hit ")
	decoder := json.NewDecoder(req.Body)
    var e Event
    err := decoder.Decode(&e)
    if err != nil {
        panic(err)
    }
    defer req.Body.Close()
    log.Println(e)
    var events = []Event{e}
    StoreEvents(events)
}

func ReadAllEventsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(ReadEvents())
}

func Serve() {
	initDatabase()
	var host = "localhost"
	var port = "8000"

	r := mux.NewRouter()
	r.HandleFunc("/calendar", CalendarHandler)
	r.HandleFunc("/plants", PlantViewHandler)
	r.HandleFunc("/light/{id}", LightHandler)
	r.HandleFunc("/pump/{id}", PumpHandler)
	r.HandleFunc("/addevent", AddEventHandler)
	r.HandleFunc("/read", ReadAllEventsHandler)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("../public"))))
	http.Handle("/", r)
	log.Printf("goLang server listening on %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
