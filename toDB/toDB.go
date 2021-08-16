package todb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var db sql.DB

func ConnectDB() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")
}
func InsertToDB(title string, rating float32, category []string, description, director string, writers, stars []string) (int64, error) {
	q := `insert into film(title,rating,category,description,director,writers,stars)
	values(?,?,?,?,?,?,?)`
	res, err := db.Exec(q, title, rating, strings.Join(category, ", "), description, director, strings.Join(writers, ", "), strings.Join(stars, ", "))
	if err != nil {
		log.Printf("Error inserting data of film %q: %v", title, err)
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error inserting data of film %q: %v", title, err)
		return 0, err
	}
	return rows, nil
}
