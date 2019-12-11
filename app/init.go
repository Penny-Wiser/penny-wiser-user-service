package app

import (
	"github.com/chenlu-chua/penny-wiser/user-service/logging"
	"github.com/chenlu-chua/penny-wiser/user-service/svccontainer"
	"log"
)

func (app *UserServiceApplication) startDependencyInjection() {
	app.diServiceContainer = &svccontainer.DIServiceContainer{Config: app.config}
	app.diServiceContainer.StartDependencyInjection()

	// Service injection
	app.server.Init(app.diServiceContainer)
	err := app.server.RegisterHandlers()
	if err != nil {
		log.Fatal(err)
	}

	logging.Logger.Info("Started all services successfully")
}
