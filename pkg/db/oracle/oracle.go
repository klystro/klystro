package oracle

import (
	"database/sql"
	"log"

	_ "github.com/godror/godror"
)

func ConnectOracleDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Connected to OracleDB!")
	return db, nil
}
