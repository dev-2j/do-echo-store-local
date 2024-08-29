package routex

import (
	"myapp/action"
	"myapp/authenx"
	"myapp/constanx"
	"myapp/typex"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.GET("/", action.XAction)
	e.POST("/login", authenx.XLogin)
	e.POST("/register", authenx.XRegister)

	r := e.Group(``)
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(typex.JW)
		},
		SigningKey: []byte(constanx.KeyDev),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("/product-get", action.XProductGet)
	r.POST("/product-create", action.XProductCreate)
	r.PUT("/product-update/:id", action.XProductUpdate)
	r.DELETE("/product-delete/:id", action.XProductDelete)

	e.Logger.Fatal(e.Start(":1323"))
}
