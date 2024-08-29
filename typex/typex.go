package typex

import "github.com/golang-jwt/jwt/v5"

type JW struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
