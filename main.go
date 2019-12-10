package main

import (
	"github.com/chenlu-chua/penny-wiser/user-service/app"
)

func main() {
	application := app.New()

	application.Init()

	application.Start("3000")
}
