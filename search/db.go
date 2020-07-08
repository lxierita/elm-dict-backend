package search

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	//justify it
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "db"
	port     = "3306"
	user     = "root"
	password = "example"
	dbname   = "test"
	sslmode  = "disable"
)

//ConnectDB to the database
func ConnectDB() {
	dsn := fmt.Sprintf("%v:%v@%v/%v", user, password, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		log.Println(err)
	}
	log.Println("Connected")
}
