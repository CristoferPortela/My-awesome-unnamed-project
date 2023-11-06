package server

import (
	"awesomeProject/handler"
	"awesomeProject/pkg"
	"awesomeProject/server/contexts"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, apiType string) {

	local := e.Group("/")
	contexts.UseAPIContext(local, false)

	api := e.Group("/api")
	contexts.UseAPIContext(api, true)

	routes := []pkg.Route{
		{Path: "", Handler: handler.Home},
	}

	pkg.RenderRoutes(routes, pkg.RoutesConfig{ApiType: apiType, Local: local, Api: api})

}
