package authenx

import (
	"myapp/constanx"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type DtoLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func GetToken(c string) (*string, error) {

	claims := &jwtCustomClaims{
		c,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(constanx.KeyDev))
	if err != nil {
		return nil, err
	}

	return &t, nil
}
