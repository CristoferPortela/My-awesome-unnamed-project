package contexts

import (
	"github.com/labstack/echo/v4"
)

type APIContext struct {
	echo.Context
	IsAPI bool
}

func UseAPIContext(e *echo.Group, isAPI bool) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &APIContext{
				c,
				isAPI,
			}
			return next(cc)
		}
	})

}
