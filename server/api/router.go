package api

import (
	"github.com/kazdevl/twimec/usecase"
	"github.com/kazdevl/twimec/web"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// web server
	e.Renderer = web.NewTemplate()
	e.GET("/", web.ContentListPage)
	e.GET("/:content_id", web.ContentPage)
	e.GET("/:content_id/:chapter_number", web.ChapterPage)

	// TODO impl
	e.POST("/config/auth", usecase.RegisterCredential)
	e.POST("/config/contents", usecase.RegisterContent)
	e.PUT("/config/contents/:user_name", usecase.UpdateContent)
	e.GET("/contents", usecase.GetContents)
	e.GET("/contents/:author_name", usecase.GetContent)
	e.GET("/contents/:author_name/:chapter", usecase.GetChapter)
}
