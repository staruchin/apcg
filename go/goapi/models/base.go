package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func init() {
	connection := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	Db, err = sql.Open("mysql", connection)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 10; i++ {
		if err = Db.Ping(); err != nil {
			time.Sleep(time.Second * 1)
			fmt.Printf("retrying... %v\n", i+1)
		} else {
			fmt.Println("Successfully connected")
			return
		}
	}
	log.Fatalln("Failed to connect")
}
