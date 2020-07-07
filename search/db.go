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
	host     = "localhost"
	port     = "3326"
	user     = "root"
	password = "example"
	dbname   = "dictionary"
	sslmode  = "disable"
)

//ConnectDB to the database
func ConnectDB() {
	dsn := fmt.Sprintf("%v:%v@/%v", user, password, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
}