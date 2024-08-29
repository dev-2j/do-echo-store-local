package authenx

import (
	"fmt"
	"myapp/connx"
	"myapp/constanx"
	"myapp/logx"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type DtoRegister struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func XRegister(c echo.Context) error {

	var rto bson.M
	if err := c.Bind(&rto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if len(rto) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	KEYPASS := constanx.KeyDev
	passkey := fmt.Sprintf("%s|%s", rto[`password`], KEYPASS)
	passEnc, err := bcrypt.GenerateFromPassword([]byte(passkey), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// set rto
	rto[`_id`] = uuid.New().String()
	rto[`password`] = string(passEnc)
	// insert

	// connext mongo
	mgx := connx.Mg()
	if err := mgx.ConnextMongo(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Connext Mongo Failed"})
	}
	// connx table and collection name
	mgx.Cnx = mgx.Client.Database("mydev").Collection("vendors")
	_, err = mgx.Cnx.InsertOne(c.Request().Context(), rto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, logx.Info("Register Success"))

}
