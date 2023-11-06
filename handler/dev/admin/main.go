package devadmin

import (
	"awesomeProject/model"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "admin-login.html", "Admin")
}

type homeData struct {
	Title  string
	Models []model.ModelMenu
}

func Home(c echo.Context) error {
	registeredModels, err := model.RegisterModel()

	if err != nil {
		log.Fatal("Error when showing models")
	}

	data := homeData{
		Title:  "Admin",
		Models: registeredModels,
	}
	return c.Render(http.StatusOK, "admin-home.html", data)
}
