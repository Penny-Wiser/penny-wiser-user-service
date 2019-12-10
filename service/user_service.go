package service

import (
	"context"
	"github.com/chenlu-chua/penny-wiser/user-service/config"
	"github.com/chenlu-chua/penny-wiser/user-service/datastore"
	"github.com/chenlu-chua/penny-wiser/user-service/model"
)

type UserService interface {
	RegisterUser(ctx context.Context) (*model.User, error)
}

type userService struct {
	config         *config.GeneralConfig
	userMongoStore *datastore.UserMongoStore
}

func NewUserService(config *config.GeneralConfig, userMongoStore *datastore.UserMongoStore) UserService {
	return &userService{
		config,
		userMongoStore,
	}
}

func (s *userService) RegisterUser(ctx context.Context) (*model.User, error) {
	return nil, nil
}
