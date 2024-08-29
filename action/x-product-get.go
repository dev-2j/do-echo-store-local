package action

import (
	"myapp/authenx"
	"myapp/connx"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func XProductGet(c echo.Context) error {
	// connext mongo
	mgx := connx.Mg()
	if err := mgx.ConnextMongo(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Connext Mongo Failed"})
	}
	// connx table and collection name
	mgx.Cnx = mgx.Client.Database("mydev").Collection("products")

	// check login
	userName, ex := authenx.GetLogin(c)
	if ex != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "กรุณาเข้าสู่ระบบ ใหม่อีกครั้ง"})
	}
	// query หา ชื่อผู้ใช้งาน
	row, err := mgx.Cnx.Find(c.Request().Context(), bson.M{"create_user": userName.ID})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Username or password is incorrect"})
	}
	var results []bson.M
	if err := row.All(c.Request().Context(), &results); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": results})
}
