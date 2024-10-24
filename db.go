package main

import (
	"database/sql"
	"log"
	_ "time"
	_ "github.com/go-sql-driver/mysql"
)

func initdb(){
	//db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
}

func dbinfo(){
	log.Println("DB drivers: ", sql.Drivers())
}
