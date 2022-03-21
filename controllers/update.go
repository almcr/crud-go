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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user_id := ctx.Param("id")
		var user models.UserData

		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// Update User in collection
		upsert := true
		defer cancel()
		r, err := database.UsersCollection.UpdateOne(
			c,
			bson.M{"user_id": user_id},
			&user,
			&options.UpdateOptions{
				Upsert: &upsert,
			},
		)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Update Data file if diff
		// this is super inneficient we need some watch utility to track of changes
		// in some field of the document collection (Change stream ?)
		data_path := models.DataPathString(user_id)
		data, err := os.ReadFile(data_path)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		if string(data) != user.Data {
			os.WriteFile(data_path, data, 0777)
		}

		ctx.JSON(http.StatusOK, r)
	}
}
