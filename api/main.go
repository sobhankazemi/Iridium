package main

import (
	"api/driver"
	"api/handlers"
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"
)

func initDB(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `create table information (
	id serial primary key,
	Os varchar(100) default ' ',
	KernelName varchar(100) default ' ',
	HostName varchar(100) default ' ',
	KernelRelease varchar(100) default ' ',
	KernelVersion varchar(100) default ' ',
	Machine varchar(100) default ' ',
	Processor varchar(100) default ' ',
	HwPlatform varchar(100) default ' ',
	UsedSpace varchar(100) default ' ',
	DateTime  varchar(100) default ' ')`
	_, err := db.ExecContext(ctx, query)
	return err
}

func main() {
	db := driver.NewDB()
	handler := handlers.New(db)
	err := initDB(db)
	if err != nil {
		log.Println(err.Error())
	}
	server := http.Server{
		Addr:    ":8080",
		Handler: GetRoutes(handler),
	}
	server.ListenAndServe()
}
