package influxdb

import (
	"context"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func ConnectInfluxDB(uri, token string) (influxdb2.Client, error) {
	client := influxdb2.NewClient(uri, token)

	// Use the client to ping InfluxDB
	_, err := client.Health(context.Background())
	if err != nil {
		return nil, err
	}

	log.Println("Connected to InfluxDB!")
	return client, nil
}
