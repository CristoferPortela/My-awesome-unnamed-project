package app

import (
	"awesomeProject/server"
	"github.com/labstack/echo/v4"
)

func StartApp() {
	e := echo.New()
	server.Router(e)
	e.Logger.Fatal(e.Start(":1324"))
}
