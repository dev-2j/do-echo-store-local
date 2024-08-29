package action

import (
	"myapp/authenx"
	"myapp/connx"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func XProductUpdate(c echo.Context) error {

	userLogin, ex := authenx.GetLogin(c)
	if ex != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "กรุณาเข้าสู่ระบบ ใหม่อีกครั้ง"})
	}

	var rto bson.M
	if err := c.Bind(&rto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if len(rto) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// connext mongo
	mgx := connx.Mg()
	if err := mgx.ConnextMongo(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Connext Mongo Failed"})
	}
	// connx table and collection name
	mgx.Cnx = mgx.Client.Database("mydev").Collection("products")

	// update item
	rto["update_user"] = userLogin.ID
	rto["updated_at"] = time.Now().Local()

	filter := bson.M{"_id": c.Param("id")}
	update := bson.M{"$set": rto}

	_, err := mgx.Cnx.UpdateOne(c.Request().Context(), filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return nil
}
