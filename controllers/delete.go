package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/almcr/crud-go/database"
	"github.com/almcr/crud-go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := ctx.Param("id")

		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// Delete User from collection
		defer cancel()
		r, err := database.UsersCollection.DeleteOne(c, bson.M{"user_id": user_id})
		if err != nil {
			log.Fatal(err)
			return
		}

		// Delete User associated file
		err = os.Remove(models.DataPathString(user_id))
		if err != nil {
			log.Fatal(err)
			return
		}

		ctx.JSON(http.StatusOK, r)
	}
}
