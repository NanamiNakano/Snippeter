package controllers

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	global "github.com/nanaminakano/snippeter/models"
	"github.com/nanaminakano/snippeter/views"
	"html/template"
	"net/http"
	"strings"
)

func Index(c echo.Context) error {
	result, err := global.Database.Snippet.Query().All(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed quering data from database: %w", err))
	}

	count, err := global.Database.Snippet.Query().Count(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed quering snippets count: %w", err))
	}

	return c.Render(http.StatusOK, "index", map[string]any{
		"Snippets": result,
		"Count":    count,
	})
}

func GetSnippet(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invaild snippet id: %w", err))
	}
	result, err := global.Database.Snippet.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed getting snippet: %w", err))
	}

	content := bluemonday.UGCPolicy().SanitizeBytes(markdown.ToHTML([]byte(result.Content), nil, global.MarkdownRenderer))

	data := views.Data{
		Content: template.HTML(content),
	}

	return c.Render(http.StatusOK, "snippet", &data)
}

func CreateSnippet(c echo.Context) error {
	var snippet views.Snippet
	if err := c.Bind(&snippet); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("bad request: %w", err))
	}

	snippet.Content = strings.ReplaceAll(snippet.Content, "\r\n", "\n")

	result, err := global.Database.Snippet.Create().
		SetID(uuid.New()).
		SetContent(snippet.Content).
		Save(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed saving data: %w", err))
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/p/%s", result.ID))
}
