package datastore

import "go.mongodb.org/mongo-driver/mongo"

const (
	billingCollectionName = "bills"
)

type BillingMongoStore struct {
	*MongoDatastore
	collection *mongo.Collection
}

func NewBillingStore(mongoDatastore *MongoDatastore) *BillingMongoStore {
	collection := mongoDatastore.DB.Collection(billingCollectionName)
	return &BillingMongoStore{
		mongoDatastore,
		collection,
	}
}
