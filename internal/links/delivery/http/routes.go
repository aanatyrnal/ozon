package http

import (
	"github.com/labstack/echo/v4"
	"ozon/internal/links"
)

func MapLinksRoutes(LinkGroup *echo.Group, h links.Handlers) {
	LinkGroup.POST("/", h.Create())
	LinkGroup.GET("/", h.Read())
}
