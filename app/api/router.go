package api

import (
	"github.com/bhill77/web/app/api/handler"
	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, handler *handler.ArticleHandler) {
	e.GET("/articles", handler.GetAll)
	e.POST("/articles", handler.Add)
}
