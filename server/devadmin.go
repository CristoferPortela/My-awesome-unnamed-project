package server

import (
	devadmin "awesomeProject/handler/dev/admin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func DevAdmin(e *echo.Echo) {
	adm := e.Group("/dev/admin")
	adm.Use(middleware.Static("/static"))
	adm.Static("/static", "template/dev/admin/assets")

	adm.GET("/home", devadmin.Home)
	adm.GET("", devadmin.Login)
}
