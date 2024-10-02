package postgres

import (
	"klystro/pkg/db/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDB(cfg config.DBConfig) *PostgresDB {
	db, err := gorm.Open(postgres.Open(cfg.URI), &gorm.Config{})
	if err != nil {
		panic(err) // Handle error appropriately
	}

	return &PostgresDB{db: db}
}

func (db *PostgresDB) Connect() error {
	return nil
}

func (db *PostgresDB) Close() error {
	sqlDB, err := db.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
