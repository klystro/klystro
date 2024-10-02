package influxdb

import (
	"klystro/pkg/db/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDB struct {
	client influxdb2.Client
}

func NewInfluxDB(cfg config.DBConfig) *InfluxDB {
	client := influxdb2.NewClient(cfg.URI, cfg.Token)
	return &InfluxDB{client: client}
}

func (db *InfluxDB) Connect() error {
	// InfluxDB's client does not require a specific connection method,
	// but you can perform a simple write to test the connection if needed.
	return nil
}

func (db *InfluxDB) Close() error {
	db.client.Close()
	return nil
}
