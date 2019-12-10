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
	r := server.NewRouter()

	logging.InitializeLogger(generalConfig)

	return &UserServiceApplication{
		config: generalConfig,
		server: r,
	}
}

// Init injects all dependencies and starts all services
func (app *UserServiceApplication) Init() error {

	return nil
}

// Start starts the http server listening on port serverPort
func (app *UserServiceApplication) Start(serverPort string) error {

	return http.ListenAndServe(":"+serverPort, app.server.GetRouter())
}
