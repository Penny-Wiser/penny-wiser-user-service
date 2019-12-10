package svccontainer

import (
	"github.com/chenlu-chua/penny-wiser/user-service/config"
	"github.com/chenlu-chua/penny-wiser/user-service/datastore"
)

type DIServiceContainer struct {
	Config *config.GeneralConfig

	BaseMongoStore *datastore.MongoDatastore

	UserMongoStore *datastore.UserMongoStore
}

func (container *DIServiceContainer) StartDependencyInjection() {

	// Init mongodb
	container.BaseMongoStore = datastore.New(&container.Config.MongoDbConfig)

	// Datastore
	container.UserMongoStore = datastore.NewUserMongoStore(container.BaseMongoStore)

	// Services

}
