package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlBookstoreUsername = "mysqlBookstoreUsername"
	mysqlBookstorePassword = "mysqlBookstorePassword"
	mysqlBookstoreHost     = "mysqlBookstoreHost"
	mysqlBookstoreDB       = "mysqlBookstoreDB"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlBookstoreUsername)
	password = os.Getenv(mysqlBookstorePassword)
	host     = os.Getenv(mysqlBookstoreHost)
	db       = os.Getenv(mysqlBookstoreDB)
)

func init() {
	var err error
	dataSources := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, db,
	)

	Client, err = sql.Open("mysql", dataSources)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully connected")
}
