package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHandler struct {
	client *mongo.Client
	db     string
}

// GetConnection func returns MongoDB's client
func GetConnection(host string, dbName string) *MongoHandler {
	ctx, cn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cn()
	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI(host))
	con := &MongoHandler{
		client: cl,
		db:     dbName,
	}
	return con
}
