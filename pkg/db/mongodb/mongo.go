package mongodb

import (
	"context"
	"klystro/pkg/db/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB(cfg config.DBConfig) *MongoDB {
	clientOptions := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err) // Handle error appropriately
	}

	return &MongoDB{client: client}
}

func (db *MongoDB) Connect() error {
	// Check the connection
	return db.client.Ping(context.Background(), nil)
}

func (db *MongoDB) Close() error {
	return db.client.Disconnect(context.Background())
}
