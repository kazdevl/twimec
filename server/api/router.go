package api

import (
	"github.com/kazdevl/twimec/usecase"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/config/auth", usecase.RegisterCredential)
	e.POST("/config/contents", usecase.RegisterContent)
	e.PUT("/config/contents/:user_name", usecase.UpdateContent)
	e.GET("/contents", usecase.GetContents)
	e.GET("/contents/:author_name", usecase.GetContent)
	e.GET("/contents/:author_name/:chapter", usecase.GetChapter)
}
