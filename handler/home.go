package handler

import (
	"awesomeProject/server/contexts"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {

	cc := c.(*contexts.APIContext)

	fmt.Println(cc)

	return c.Render(http.StatusOK, "home.html", "")
}
