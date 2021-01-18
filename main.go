package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./logic"

	"./db"
	_ "github.com/mattn/go-sqlite3"
)

type indicator struct {
	Main  db.Rectangle   `json:"main"`
	Input []db.Rectangle `json:"input"`
}

func filteringServer(w http.ResponseWriter, r *http.Request) {
	var i indicator
	if r.Method == "GET" {
		db.GetAll(w)
	} else if r.Method == "POST" {

		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &i)
		if err != nil {
			println(err)
		}
		res := logic.FindOverlap(i.Main, i.Input)
		db.AddToDb(res)
	} else {
		w.Write([]byte("Wrong Method"))
	}

}

func main() {
	database, _ := sql.Open("sqlite3", "./db.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS rectangle (x INTEGER,y INTEGER,width INTEGER,height INTEGER,date TEXT)")
	statement.Exec()
	http.HandleFunc("/", filteringServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
