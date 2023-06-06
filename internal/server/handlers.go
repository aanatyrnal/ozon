package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	linksHttp "ozon/internal/links/delivery/http"
	linksRepository "ozon/internal/links/repository"
	linksServices "ozon/internal/links/service"
	linksUseCase "ozon/internal/links/usecase"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))

	linksRepo := linksRepository.NewLinksRepository(s.db)

	linksServices := linksServices.NewLinksService(s.cfg, linksRepo)

	linksUS := linksUseCase.NewLinksUseCase(s.cfg, linksServices)

	linksHandlers := linksHttp.NewLinksHandlers(s.cfg, linksUS)
	v1 := e.Group("/api")

	linksGroup := v1.Group("")

	linksHttp.MapLinksRoutes(linksGroup, linksHandlers)

	return nil

}
