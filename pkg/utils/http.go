package utils

import (
	"context"

	"github.com/labstack/echo/v4"
)

func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}
