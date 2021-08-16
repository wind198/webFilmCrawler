package todb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "film",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println(db, *db)
	fmt.Println("Connected")
	return db
}

func InsertToDB(db *sql.DB, title string, rating float32, category []string, description, director string, writers, stars []string) (int64, error) {
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

func CreateTable(db *sql.DB) {
	dropTableQuery := `drop table if exists film;`
	createTableQuery := `
	create table film(
		id int NOT NULL AUTO_INCREMENT,
			title     varchar(255)  ,
			rating      float(5, 2),
			category    varchar(255),
			description LONGTEXT,
			director    varchar(255),
			writers     varchar(255),
			stars       varchar(255),
			 PRIMARY KEY (id)
		);`
	log.Println(db)
	if _, err := db.Exec(dropTableQuery); err != nil {
		log.Println(err)
	}
	if _, err := db.Exec(createTableQuery); err != nil {
		log.Println(err)
	}

}
