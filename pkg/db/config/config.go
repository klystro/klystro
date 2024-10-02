package config

import (
	"klystro/config"
	"log"
)

type DBConfig struct {
	URI         string
	Token       string // For InfluxDB
	ConnectFlag bool
}

func GetDBConfig(dbType string) DBConfig {
	switch dbType {
	case "MONGODB":
		return DBConfig{
			URI:         config.AppConfig.MongoDBURI,
			ConnectFlag: config.AppConfig.MongoDBConnect,
		}
	case "MYSQL":
		return DBConfig{
			URI:         config.AppConfig.MySQLURI,
			ConnectFlag: config.AppConfig.MySQLConnect,
		}
	case "POSTGRES":
		return DBConfig{
			URI:         config.AppConfig.PostgresURI,
			ConnectFlag: config.AppConfig.PostgresConnect,
		}
	case "INFLUXDB":
		return DBConfig{
			URI:         config.AppConfig.InfluxDBURI,
			Token:       config.AppConfig.InfluxDBToken,
			ConnectFlag: config.AppConfig.InfluxDBConnect,
		}
	case "ORACLEDB":
		return DBConfig{
			URI:         config.AppConfig.OracleDBURI,
			ConnectFlag: config.AppConfig.OracleDBConnect,
		}
	default:
		log.Fatalf("unknown database type: %v", dbType)
		return DBConfig{}
	}
}
