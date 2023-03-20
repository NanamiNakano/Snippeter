package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/nanaminakano/snippeter/controllers"
	global "github.com/nanaminakano/snippeter/models"
	"github.com/nanaminakano/snippeter/models/ent"
	"github.com/nanaminakano/snippeter/views"
	"html/template"
	"log"
	"net/http"
)

func main() {
	var err error
	global.Database, err = ent.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to the database: %v", err)
	}
	defer func() {
		_ = global.Database.Close()
	}()
	if err := global.Database.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	t := views.Template{
		Templates: template.Must(template.ParseFS(views.TemplateFS, "template/*.tmpl")),
	}
	e := echo.New()
	e.Renderer = &t
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if httpError, ok := err.(*echo.HTTPError); ok {
			_ = c.JSON(httpError.Code, httpError.Error())
		} else {
			_ = c.JSON(http.StatusInternalServerError, err)
		}

		log.Println(err)
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.GET("/", controllers.Index)
	e.POST("/", controllers.CreateSnippet)
	e.GET("/p/:id", controllers.GetSnippet)
	e.Logger.Fatal(e.Start(":8080"))
}
