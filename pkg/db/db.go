package db

import (
	"context"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"klystro/config"
	"klystro/pkg/db/mongodb"
)

var (
	MongoClient  *mongo.Client
	MySQLDB      *gorm.DB
	PostgresDB   *gorm.DB
	OracleDB     *gorm.DB
	InfluxClient influxdb2.Client
)

func InitDatabases() {
	var err error

	// MongoDB
	MongoClient, err = mongodb.ConnectMongoDB(config.GetOrThrow("MONGODB_URI"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// MySQL
	MySQLDB, err = gorm.Open(mysql.Open(config.GetOrThrow("MYSQL_URI")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// PostgreSQL
	PostgresDB, err = gorm.Open(postgres.Open(config.GetOrThrow("POSTGRES_URI")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// // InfluxDB
	// InfluxClient, err = influxdb.ConnectInfluxDB(config.GetOrThrow("INFLUXDB_URI"), config.GetOrThrow("INFLUXDB_TOKEN"))
	// if err != nil {
	// 	log.Fatalf("Failed to connect to InfluxDB: %v", err)
	// }
}

func CloseDatabases() {
	if err := MongoClient.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect MongoDB: %v", err)
	}

	// Close MySQL
	if sqlDB, err := MySQLDB.DB(); err == nil {
		sqlDB.Close()
	}

	// Close PostgreSQL
	if sqlDB, err := PostgresDB.DB(); err == nil {
		sqlDB.Close()
	}

	// // Close OracleDB
	// if sqlDB, err := OracleDB.DB(); err == nil {
	// 	sqlDB.Close()
	// }

	// InfluxClient.Close()
}
