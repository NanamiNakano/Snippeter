package views

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Snippet struct {
	Language string `form:"language"`
	Content  string `form:"content"`
}

type Data struct {
	Content template.HTML
}

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(writer io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(writer, name, data)
}
