package server

import (
	"awesomeProject/cmd"
	"awesomeProject/model"
	"awesomeProject/pkg"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func configureTemplatePaths(e *echo.Echo, db *gorm.DB) {
	var tpl model.Setting
	db.Where("setting_name = ?", []string{"template"}).First(&tpl)

	if tpl.ID == 0 {
		cmd.InitializeDB(db)
		configureTemplatePaths(e, db)
	}

	var files []string

	files = append(files, pkg.GetHtmlFrom("template/dev/admin")...)
	files = append(files, pkg.GetHtmlFrom(fmt.Sprintf("template/%s", tpl.SettingValue))...)

	t := &Template{
		templates: template.Must(template.ParseFiles(files...)),
	}
	e.Renderer = t

}
