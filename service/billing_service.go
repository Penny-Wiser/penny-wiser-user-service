package service

import (
	"github.com/chenlu-chua/penny-wiser/user-service/config"
	"github.com/chenlu-chua/penny-wiser/user-service/datastore"
)

type BillingService interface {
}

type billingService struct {
	config            *config.GeneralConfig
	billingMongoStore *datastore.BillingMongoStore
}

func NewBillingService(config *config.GeneralConfig, billingMongoStore *datastore.BillingMongoStore) BillingService {
	return &billingService{
		config,
		billingMongoStore,
	}
}
