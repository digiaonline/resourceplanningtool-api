package modules

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(dbtype string, dbaddress string) {
	var err error
	db, err = sqlx.Connect(dbtype, dbaddress)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
