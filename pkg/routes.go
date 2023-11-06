package pkg

import (
	"awesomeProject/handler"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Path    string
	Handler echo.HandlerFunc
}

type RoutesConfig struct {
	ApiType string
	Local   *echo.Group
	Api     *echo.Group
}

// RenderRoutes will ge all routes passed to it and render it to local and/or api, if apiType is not configured, it will throw an error on every route
func RenderRoutes(r []Route, rc RoutesConfig) {
	if rc.ApiType == WITH_API || rc.ApiType == NO_API || rc.ApiType == ONLY_API {
		for _, route := range r {
			if rc.ApiType == WITH_API || rc.ApiType == NO_API {
				rc.Local.GET(route.Path, handler.Home)
			}

			if rc.ApiType == WITH_API || rc.ApiType == ONLY_API {
				rc.Api.GET("/Api"+route.Path, handler.Home)
			}
		}
	} else {
		rc.Local.Any("", handler.HandleApiTypeNotFound)
		rc.Local.Any("*", handler.HandleApiTypeNotFound)
		rc.Api.Any("", handler.HandleApiTypeNotFound)
		rc.Api.Any("/*", handler.HandleApiTypeNotFound)
	}
}
