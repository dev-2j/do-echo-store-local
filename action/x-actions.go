package action

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func XAction(c echo.Context) error {
	/// test
	//git config --global user.email "you@example.com"
  git config --global user.name "Your Name"
	return c.JSON(http.StatusOK, echo.Map{"message": "Hello, World!"})

}
