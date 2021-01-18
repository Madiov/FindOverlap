package db

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

//Rectangle includs a rectangle properties.
type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
	Date   string
}

//AddToDb adds the rectangle to the database.
func AddToDb(r Rectangle) {
	r.Date = time.Now().Format("2006.01.02 15:04:05")
	database, _ := sql.Open("sqlite3", "./db.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS rectangle (x INTEGER,y INTEGER,width INTEGER,height INTEGER,date TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO rectangle (x,y,width,height,date) VALUES ( ?,?,?,?,?)")
	statement.Exec(r.X, r.Y, r.Width, r.Height, r.Date)

}

//GetAll shows all of the stored rectangles.
func GetAll(w http.ResponseWriter) {
	database, _ := sql.Open("sqlite3", "./db.db")
	rows, _ := database.Query("SELECT x,y,width,height,date FROM rectangle")
	var x, y, width, height int
	var date string
	for rows.Next() {
		rows.Scan(&x, &y, &width, &height, &date)
		re := Rectangle{X: x, Y: y, Width: width, Height: height, Date: date}
		json.NewEncoder(w).Encode(re)

	}
}
