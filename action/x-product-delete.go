package action

import (
	"myapp/authenx"
	"myapp/connx"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func XProductDelete(c echo.Context) error {

	userLogin, ex := authenx.GetLogin(c)
	if ex != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "กรุณาเข้าสู่ระบบ ใหม่อีกครั้ง"})
	}

	if c.Param("id") == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// connext mongo
	mgx := connx.Mg()
	if err := mgx.ConnextMongo(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Connext Mongo Failed"})
	}
	// connx table and collection name
	mgx.Cnx = mgx.Client.Database("mydev").Collection("products")

	// ลบ สินค้า
	_, err := mgx.Cnx.DeleteOne(c.Request().Context(), bson.M{"_id": c.Param("id"), "create_user": userLogin.ID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return nil
}
