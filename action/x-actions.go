package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func XAction(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "Hello, World!"})
}
