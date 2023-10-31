package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	message := HelloMessage()
	return c.String(http.StatusOK, message)
}
