package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	MongoDBURI    string
	MySQLURI      string
	PostgresURI   string
	OracleDBURI   string
	InfluxDBURI   string
	InfluxDBToken string
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = &Config{
		MongoDBURI:    GetOrThrow("MONGODB_URI"),
		MySQLURI:      GetOrThrow("MYSQL_URI"),
		PostgresURI:   GetOrThrow("POSTGRES_URI"),
		OracleDBURI:   GetOrThrow("ORACLEDB_URI"),
		InfluxDBURI:   GetOrThrow("INFLUXDB_URI"),
		InfluxDBToken: GetOrThrow("INFLUXDB_TOKEN"),
	}
}

func GetOrThrow(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}
