package datastore

import (
	"context"
	"github.com/chenlu-chua/penny-wiser/user-service/model"
	"github.com/kelseyhightower/confd/log"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollectionName = "users"
)

type UserMongoStore struct {
	*MongoDatastore
	collection *mongo.Collection
}

func NewUserMongoStore(mongoDataStore *MongoDatastore) *UserMongoStore {
	collection := mongoDataStore.DB.Collection(userCollectionName)
	return &UserMongoStore{
		mongoDataStore,
		collection,
	}
}

func (s *UserMongoStore) Create(ctx context.Context, userModel *model.User) (*model.User, error) {

	res, err := s.collection.InsertOne(ctx, userModel)
	if err != nil {
		return nil, err
	}

	var id string
	if objId, ok := res.InsertedID.(primitive.ObjectID); ok {
		id = objId.Hex()
	} else {
		log.Info("Failed to create id from objectID")
		id = ""
	}

	userAppendId := userModel
	userAppendId.Id = id
	return userAppendId, nil
}

func (s *UserMongoStore) Find(ctx context.Context, criteria interface{}) (*model.User, error) {
	var user model.User
	err := s.collection.FindOne(ctx, criteria).Decode(&user)
	if err != nil {
		return nil, err
	}

	user.Id = user.RawId.Hex()
	user.RawId = nil

	return &user, nil
}

func (s *UserMongoStore) Update(ctx context.Context, user *model.User) (*model.User, error) {
	criteria := M{"id": user.Id}
	persistedUser, err := s.Find(ctx, criteria)
	if err != nil {
		return nil, err
	}

	res, err := s.collection.ReplaceOne(ctx, M{"id": persistedUser.Id, "updated_at": persistedUser.UpdatedAt}, user)
	if err != nil {
		return nil, err
	}

	if res.ModifiedCount < 1 {

	}

	return nil, nil
}
