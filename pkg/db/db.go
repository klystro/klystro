package db

import (
	"fmt"
	"log"
	"sync"

	"klystro/pkg/db/config"
	"klystro/pkg/db/influxdb"
	"klystro/pkg/db/mongodb"
	"klystro/pkg/db/mysql"
	"klystro/pkg/db/oracle"
	"klystro/pkg/db/postgres"
)

type Database interface {
	Connect() error
	Close() error
}

type DBType string

const (
	Mongo    DBType = "MONGODB"
	MySQL    DBType = "MYSQL"
	Postgres DBType = "POSTGRES"
	Influx   DBType = "INFLUXDB"
	Oracle   DBType = "ORACLEDB"
)

var (
	dbOnce     sync.Once
	dbInstance = make(map[DBType]Database)
)

func NewDatabase(dbType DBType) (Database, error) {
	cfg := config.GetDBConfig(string(dbType))

	if !cfg.ConnectFlag {
		log.Printf("Connection flag is false for %v; skipping connection.", dbType)
		return nil, nil
	}

	// Ensure only one instance is created for each database type
	var err error
	dbOnce.Do(func() {
		switch dbType {
		case Mongo:
			dbInstance[Mongo] = mongodb.NewMongoDB(cfg)
			err = dbInstance[Mongo].Connect()
		case MySQL:
			dbInstance[MySQL] = mysql.NewMySQLDB(cfg)
			err = dbInstance[MySQL].Connect()
		case Postgres:
			dbInstance[Postgres] = postgres.NewPostgresDB(cfg)
			err = dbInstance[Postgres].Connect()
		case Influx:
			dbInstance[Influx] = influxdb.NewInfluxDB(cfg)
			err = dbInstance[Influx].Connect()
		case Oracle:
			dbInstance[Oracle] = oracle.NewOracleDB(cfg)
			err = dbInstance[Oracle].Connect()
		default:
			err = fmt.Errorf("unknown database type: %v", dbType)
		}
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to %v: %v", dbType, err)
	}

	return dbInstance[dbType], nil
}

func InitDatabases() {
	dbTypes := []DBType{Mongo, MySQL, Postgres, Influx, Oracle}

	for _, dbType := range dbTypes {
		db, err := NewDatabase(dbType) // Declare err here
		if err != nil {
			log.Printf("Failed to create %v client: %v", dbType, err)
		} else if db != nil {
			log.Printf("Successfully connected to %v", dbType)
			dbInstance[dbType] = db
		}
	}
}

func CloseDatabases() {
	for _, db := range dbInstance {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Printf("Failed to disconnect: %v", err)
			}
		}
	}
}
