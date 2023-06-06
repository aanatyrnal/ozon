package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"ozon/config"
	"ozon/internal/links"
	"ozon/internal/models"
	"ozon/pkg/utils"
)

type linksHandlers struct {
	cfg     *config.Config
	linksUC links.UseCase
}

type Response struct {
	Status string `json:"status"`
	Text   string `json:"error"`
}

func NewLinksHandlers(cfg *config.Config, linksUS links.UseCase) links.Handlers {
	return &linksHandlers{cfg: cfg, linksUC: linksUS}
}

func (h *linksHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "links.Create")
		defer span.Finish()

		linksC := &models.Links{}

		if err := utils.ReadRequest(c, linksC); err != nil {
			fmt.Println(err)
			return c.JSON(404, Response{Status: "404", Text: "Bad Data"})
		}
		fmt.Println("span, ctx", span, ctx)
		createdLinks, err := h.linksUC.Create(ctx, linksC)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusConflict, Response{
				Status: "error",
				Text:   "link Already Exists",
			})
		}

		return c.JSON(http.StatusCreated, createdLinks)
	}
}

func (h *linksHandlers) Read() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "linksHandler.Read")
		defer span.Finish()

		shortLink := c.QueryParam("short_link")

		linksList, err := h.linksUC.Read(ctx, shortLink)
		if err != nil {
			return c.JSON(http.StatusNotFound, Response{
				Status: "error",
				Text:   "Short link not found",
			})
		}

		return c.JSON(http.StatusOK, linksList)
	}
}
