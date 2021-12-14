package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/kazdevl/twimec/domain"
	"github.com/kazdevl/twimec/repository"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	userHomeDir, _ := os.UserHomeDir()
	return &Template{
		templates: template.Must(template.ParseGlob(fmt.Sprintf("%s/twimec/web/template/*.html", userHomeDir))),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Web struct {
	configRepo      repository.ConfigRepository
	contentinfoRepo repository.ContentInfoRepository
	chapterRepo     repository.ChapterRepository
}

func NewWeb(
	cgRepo repository.ConfigRepository,
	ciRepo repository.ContentInfoRepository,
	crRepo repository.ChapterRepository,
) *Web {
	return &Web{
		configRepo:      cgRepo,
		contentinfoRepo: ciRepo,
		chapterRepo:     crRepo,
	}
}

func (w *Web) ContentListPage(c echo.Context) error {
	contentinfoList := w.contentinfoRepo.All()
	return c.Render(http.StatusOK, "contentList", contentinfoList)
}

func (w *Web) ContentPage(c echo.Context) error {
	contentID := c.Param("content_id")
	chapters := w.chapterRepo.GetList(contentID)
	var header string
	if len(chapters) != 0 {
		pages := strings.Split(chapters[0].Pages, ",")
		header = pages[0]
	}
	type data struct {
		Header     string
		AuthorName string
		Chapters   []domain.Chapter
	}
	d := data{
		Header:     header,
		AuthorName: w.contentinfoRepo.GetAuthorName(contentID),
		Chapters:   w.chapterRepo.GetList(contentID),
	}
	return c.Render(http.StatusOK, "content", d)
}

func (w *Web) ChapterPage(c echo.Context) error {
	contentID := c.Param("content_id")
	index, _ := strconv.Atoi(c.Param("chapter_number"))
	config := w.configRepo.Get(contentID)
	latestIndex := config.LatestChapter - 1
	var previous, next int
	if index > 0 {
		previous = index - 1
	}
	if latestIndex > index {
		next = index + 1
	}

	type data struct {
		ID       string
		Previous int
		Next     int
		Pages    domain.Pages
	}
	d := data{
		ID:       contentID,
		Previous: previous,
		Next:     next,
		Pages:    w.chapterRepo.GetPages(contentID, index),
	}
	return c.Render(http.StatusOK, "chapter", d)
}
