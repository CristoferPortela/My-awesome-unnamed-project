package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleApiTypeNotFound(c echo.Context) error {
	return c.String(http.StatusBadRequest, "Database not configured successfully")
}
