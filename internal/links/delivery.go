package links

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Create() echo.HandlerFunc
	Read() echo.HandlerFunc
}
