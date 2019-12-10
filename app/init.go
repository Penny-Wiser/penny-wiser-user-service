package app

import "github.com/chenlu-chua/penny-wiser/user-service/svccontainer"

func (app *UserServiceApplication) StartDependencyInjection() {
	app.diServiceContainer = &svccontainer.DIServiceContainer{Config: app.config}
	app.diServiceContainer.StartDependencyInjection()
}
