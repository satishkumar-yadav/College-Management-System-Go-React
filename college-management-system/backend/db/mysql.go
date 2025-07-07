package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var MySQLDB *sql.DB

func InitMySQL() {
	var err error
	MySQLDB, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/college")
	if err != nil {
		log.Fatal("MySQL connect error:", err)
	}
	fmt.Println("MySql Connected")
}
