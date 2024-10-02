package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	MongoDBURI      string
	MongoDBConnect  bool
	MySQLURI        string
	MySQLConnect    bool
	PostgresURI     string
	PostgresConnect bool
	OracleDBURI     string
	OracleDBConnect bool
	InfluxDBURI     string
	InfluxDBToken   string
	InfluxDBConnect bool
}

func InitConfig() {
	viper.AutomaticEnv()

	AppConfig = &Config{
		MongoDBURI:      GetOrThrow("MONGODB_URI"),
		MongoDBConnect:  viper.GetBool("CONNECT_MONGODB"),
		MySQLURI:        GetOrThrow("MYSQL_URI"),
		MySQLConnect:    viper.GetBool("MYSQL_CONNECT"),
		PostgresURI:     GetOrThrow("POSTGRES_URI"),
		PostgresConnect: viper.GetBool("POSTGRES_CONNECT"),
		OracleDBURI:     GetOrThrow("ORACLEDB_URI"),
		OracleDBConnect: viper.GetBool("ORACLEDB_CONNECT"),
		InfluxDBURI:     GetOrThrow("INFLUXDB_URI"),
		InfluxDBToken:   GetOrThrow("INFLUXDB_TOKEN"),
		InfluxDBConnect: viper.GetBool("INFLUXDB_CONNECT"),
	}
}

func GetOrThrow(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}
