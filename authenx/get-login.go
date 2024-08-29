package authenx

import (
	"myapp/typex"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type DtoGetLogin struct {
	Name string
	ID   string
}

func GetLogin(c echo.Context) (*DtoGetLogin, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*typex.JW)

	if claims == nil {
		return nil, echo.ErrBadGateway

	}

	sx := strings.Split(claims.Name, "|")
	rtx := DtoGetLogin{
		Name: sx[0],
		ID:   sx[1],
	}
	return &rtx, nil
}
