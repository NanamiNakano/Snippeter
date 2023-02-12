package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/nanaminakano/snippeter/ent"
)

var client *ent.Client

//go:embed template/*.tmpl
var templateFS embed.FS

type Template struct {
	templates *template.Template
}

func (t *Template) Render(writer io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(writer, name, data)
}

func main() {
	var err error
	client, err = ent.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to the database: %v", err)
	}
	defer func() {
		_ = client.Close()
	}()
	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	t := Template{
		templates: template.Must(template.ParseFS(templateFS, "template/*.tmpl")),
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
	e.GET("/", index)
	e.POST("/", createSnippet)
	e.GET("/p/:id", getSnippet)
	e.Logger.Fatal(e.Start(":80"))
}

func index(c echo.Context) error {
	result, err := client.Snippet.Query().All(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed quering data from database: %w", err))
	}
	return c.Render(http.StatusOK, "index", map[string]any{
		"Snippets": result,
	})
}

type Snippet struct {
	Language string `form:"language"`
	Content  string `form:"content"`
}

func createSnippet(c echo.Context) error {
	var snippet Snippet
	if err := c.Bind(&snippet); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("bad request: %w", err))
	}
	fmt.Println(snippet)
	result, err := client.Snippet.Create().
		SetID(uuid.New()).
		SetLanguage(snippet.Language).
		SetContent(snippet.Content).
		Save(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed saving data: %w", err))
	}
	fmt.Println(result)
	return c.Redirect(http.StatusFound, fmt.Sprintf("/p/%s", result.ID))
}

func getSnippet(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invaild snippet id: %w", err))
	}
	result, err := client.Snippet.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed getting snippet: %w", err))
	}
	return c.Render(http.StatusOK, "snippet", map[string]any{
		"Snippet": result,
	})
}
