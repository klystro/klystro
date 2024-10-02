package oracle

import (
	"database/sql"
	"klystro/pkg/db/config"

	_ "github.com/godror/godror"
)

type OracleDB struct {
	db *sql.DB
}

func NewOracleDB(cfg config.DBConfig) *OracleDB {
	db, err := sql.Open("godror", cfg.URI)
	if err != nil {
		panic(err) // Handle error appropriately
	}

	// Configure connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(0)

	return &OracleDB{db: db}
}

func (db *OracleDB) Connect() error {
	// You can perform a simple query to check the connection
	return nil
}

func (db *OracleDB) Close() error {
	return db.db.Close()
}
