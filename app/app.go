package app

import (
	"github.com/chenlu-chua/penny-wiser/user-service/config"
	"github.com/chenlu-chua/penny-wiser/user-service/logging"
	"github.com/chenlu-chua/penny-wiser/user-service/server"
	"github.com/chenlu-chua/penny-wiser/user-service/svccontainer"
	"net/http"
)

// AppInterface is the main application which wraps object for the application
type AppInterface interface {
	Init() error
	Start(serverPort string) error
}

type UserServiceApplication struct {
	config             *config.GeneralConfig
	diServiceContainer *svccontainer.DIServiceContainer
	server             server.Server
}

func New() AppInterface {
	generalConfig := config.LoadConfig()
	r := server.New()
	logging.InitializeLogger(generalConfig)

	return &UserServiceApplication{
		config: generalConfig,
		server: r,
	}
}

func (app *UserServiceApplication) Init() error {
	logging.Logger.Info("Started dependency injection...")
	app.startDependencyInjection()
	return nil
}

// Start starts the http server listening on port serverPort
func (app *UserServiceApplication) Start(serverPort string) error {
	logging.Logger.Infof("Started user service on port %s", serverPort)
	return http.ListenAndServe(":"+serverPort, app.server.GetServer())
}
