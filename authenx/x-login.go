package authenx

import (
	"fmt"
	"myapp/connx"
	"myapp/constanx"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func XLogin(c echo.Context) error {

	var rto bson.M
	if err := c.Bind(&rto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	username, usernameOk := rto["user_name"].(string)
	password, passwordOk := rto["password"].(string)

	if !usernameOk || !passwordOk {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username and password are required"})
	}

	// connext mongo
	mgx := connx.Mg()
	if err := mgx.ConnextMongo(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Connext Mongo Failed"})
	}
	// connx table and collection name
	mgx.Cnx = mgx.Client.Database("mydev").Collection("vendors")

	// query หา ชื่อผู้ใช้งาน
	err := mgx.Cnx.FindOne(c.Request().Context(), bson.M{"user_name": username}).Decode(&rto)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Username or password is incorrect"})
	}

	// get password old
	storedPassword, ok := rto["password"].(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User password not found"})
	}

	// validdate password
	passkeyOld := fmt.Sprintf("%s|%s", password, constanx.KeyDev)
	if !CheckPassword(storedPassword, passkeyOld) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Username or password is incorrect"})
	}

	// สร้าง token
	idx, ok := rto["_id"].(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "id found"})
	}
	keyTokenName := fmt.Sprintf("%s|%s", username, idx)
	t, err := GetToken(keyTokenName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"username": username, "token": *t})
}
