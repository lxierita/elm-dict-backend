package search 

import (
	"fmt"
	"log"
	"context"
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
func ConnectDB(){
	dsn := fmt.Sprintf("%v:%v@/%v", user, password, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	ping()
}

func ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}