package modules

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dbtype string, dbaddress string) {
	db, err := sql.Open(dbtype, dbaddress)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
