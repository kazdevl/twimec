package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	templateDir, _ := filepath.Abs("./template")
	return &Template{
		templates: template.Must(template.ParseGlob(fmt.Sprintf("%s/*.html", templateDir))),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func ContentListPage(c echo.Context) error {
	return c.Render(http.StatusOK, "contentList", "World")
}

func ContentPage(c echo.Context) error {
	return c.Render(http.StatusOK, "content", "World")
}

func ChapterPage(c echo.Context) error {
	return c.Render(http.StatusOK, "chapter", "World")
}
