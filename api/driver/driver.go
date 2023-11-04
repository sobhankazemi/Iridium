package driver

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	connectionString := "host=iridium_db port=5432 user=admin password=1234 dbname=iridium sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
	}
	return db
}
