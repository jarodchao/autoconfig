package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"autoconfig/type"
)

var API_DB *sql.DB

func InitDB(d autoconfig.Database) {

	url := d.Dns()

	var err error

	API_DB, err = sql.Open("mysql", url)

	if err != nil {
		panic(err.Error())
	}

	API_DB.SetMaxIdleConns(d.MaxIdle)
	API_DB.SetMaxOpenConns(d.MaxOpen)
	err = API_DB.Ping()

	if err != nil {
		panic(err.Error())
	}
}