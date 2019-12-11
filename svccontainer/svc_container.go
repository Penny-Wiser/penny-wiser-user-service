package svccontainer

import (
	"github.com/chenlu-chua/penny-wiser/user-service/config"
	"github.com/chenlu-chua/penny-wiser/user-service/datastore"
	"github.com/chenlu-chua/penny-wiser/user-service/service"
)

type DIServiceContainer struct {
	Config *config.GeneralConfig

	UserService    service.UserService
	BillingService service.BillingService
}

// Initialize and start up dependencies here
func (container *DIServiceContainer) StartDependencyInjection() {

	// Init mongodb
	baseMongoStore := datastore.New(&container.Config.MongoDbConfig)

	// Datastore
	userMongoStore := datastore.NewUserMongoStore(baseMongoStore)
	billingMongoStore := datastore.NewBillingStore(baseMongoStore)

	// Service
	container.UserService = service.NewUserService(container.Config, userMongoStore)
	container.BillingService = service.NewBillingService(container.Config, billingMongoStore)
}
