package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/almcr/crud-go/database"
	"github.com/almcr/crud-go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func WriteUserDataFile(filePath string, data []byte) {
	err := os.WriteFile(filePath, data, 0777)
	if err != nil {
		// todo ho
		return
	}
}

func AddUser(user models.UserData) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// discard already inserted users
	count, err := database.UsersCollection.CountDocuments(ctx, bson.M{"user_id": user.Id})
	defer cancel()
	if err != nil {
		log.Fatal(err)
		return
	}

	if count > 0 {
		log.Printf("discard user: %s\n", user.Id)
		return
	}

	// Hash user password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Password = string(hashed)

	// Insert user data into collection
	_, err = database.UsersCollection.InsertOne(ctx, &user)
	defer cancel()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func AddUsers() gin.HandlerFunc {
	var Wg sync.WaitGroup
	// error_chan := make(chan error)

	return func(ginCtx *gin.Context) {
		decoder := json.NewDecoder(ginCtx.Request.Body)
		// open bracket [ decode
		_, err := decoder.Token()
		if err != nil {
			log.Print(err)
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Malformed user data"})
			return
		}

		// stream decoding each value aka document
		for decoder.More() {
			var user models.UserData
			err := decoder.Decode(&user)
			if err != nil {
				// ignore malformed user data
				log.Print(err)
				continue
			}

			// Add user to collection
			Wg.Add(1)
			go func() {
				defer Wg.Done()
				AddUser(user)
			}()

			// Create a user data file
			Wg.Add(1)
			go func() {
				defer Wg.Done()
				WriteUserDataFile(models.DataPathString(user.Id), []byte(user.Data))
			}()
		}
		// Wait forked goroutines
		Wg.Wait()

		// closing bracket ] decode
		_, err = decoder.Token()
		if err != nil {
			log.Print(err)
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Malformed user data"})
		}

		ginCtx.JSON(http.StatusOK, nil)
	}
}
