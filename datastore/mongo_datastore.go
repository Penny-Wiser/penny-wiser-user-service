package datastore

import (
	"context"
	"fmt"
	"github.com/chenlu-chua/penny-wiser/user-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type M map[string]interface{}

// Represents the database object for each datastore
type MongoDatastore struct {
	Config *config.MongoDbConfig
	DB     *mongo.Database
	Client *mongo.Client
}

func New(mongoConfig *config.MongoDbConfig) *MongoDatastore {
	connectionUri := connectionString(mongoConfig)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionUri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(mongoConfig.DatabaseName)

	return &MongoDatastore{
		mongoConfig,
		db,
		client,
	}
}

func connectionString(config *config.MongoDbConfig) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Username, config.Password,
		config.DatabaseHost, config.DatabasePort)
}
