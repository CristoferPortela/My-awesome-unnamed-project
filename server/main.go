package server

import (
	"awesomeProject/internal"
	"awesomeProject/model"
	"awesomeProject/server/contexts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func Router(e *echo.Echo) {
	addons := internal.ConfigureAddons()

	db, err := internal.Connection()

	if err != nil {
		log.Fatal("Couldn't initialize db")
	}

	// Defining configuration context
	var apiType model.Setting
	db.Where("setting_name = ?", []string{"api_type"}).First(&apiType)
	contexts.ConfigContext(e, db, apiType.SettingValue)

	e.Pre(middleware.RemoveTrailingSlash())
	configureTemplatePaths(e, db)

	DevAdmin(e)
	//Admin(e)
	Routes(e, apiType.SettingValue)

	for _, route := range addons.Routes {
		route(e)
	}
}
