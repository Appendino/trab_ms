package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/appendino/trab_ms/issue_service/config"
)

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.MYSQL_USER,
		config.MYSQL_PASSWORD, config.MYSQL_SERVER, config.MYSQL_DATABASE))
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
