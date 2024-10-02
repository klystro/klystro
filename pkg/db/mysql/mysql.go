package mysql

import (
	"klystro/pkg/db/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	db *gorm.DB
}

func NewMySQLDB(cfg config.DBConfig) *MySQL {
	db, err := gorm.Open(mysql.Open(cfg.URI), &gorm.Config{})
	if err != nil {
		panic(err) // Handle error appropriately
	}

	return &MySQL{db: db}
}

func (db *MySQL) Connect() error {
	return nil
}

func (db *MySQL) Close() error {
	sqlDB, err := db.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
