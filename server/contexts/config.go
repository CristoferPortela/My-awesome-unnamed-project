package contexts

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ConfigurationContext struct {
	echo.Context
	DB      *gorm.DB
	ApiType string
}

func ConfigContext(e *echo.Echo, db *gorm.DB, apiType string) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ConfigurationContext{
				c,
				db,
				apiType,
			}
			return next(cc)
		}
	})

}
