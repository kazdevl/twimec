package api

import (
	"github.com/kazdevl/twimec/repository"
	"github.com/kazdevl/twimec/usecase"
	"github.com/kazdevl/twimec/web"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, configRepo repository.ConfigRepository, contentinfoRepo repository.ContentInfoRepository, chapterRepo repository.ChapterRepository) {
	// web server
	e.Renderer = web.NewTemplate()
	w := web.NewWeb(configRepo, contentinfoRepo, chapterRepo)
	e.GET("/", w.ContentListPage)
	e.GET("/:content_id", w.ContentPage)
	e.GET("/:content_id/:chapter_number", w.ChapterPage)

	// TODO impl
	e.POST("/config/auth", usecase.RegisterCredential)
	e.POST("/config/contents", usecase.RegisterContent)
	e.PUT("/config/contents/:user_name", usecase.UpdateContent)
	e.GET("/contents", usecase.GetContents)
	e.GET("/contents/:author_name", usecase.GetContent)
	e.GET("/contents/:author_name/:chapter", usecase.GetChapter)
}
