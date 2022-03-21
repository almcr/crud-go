package app

import (
	"log"
	"os"

	"github.com/almcr/crud-go/controllers"
	"github.com/almcr/crud-go/database"
	"github.com/almcr/crud-go/middlewares"
	"github.com/almcr/crud-go/models"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Address string
	Port    string
}

func NewServer() Server {
	var server Server

	server.Port = os.Getenv("APP_PORT")
	server.Address = os.Getenv("ADDRESS")
	models.UserDataFilePath = os.Getenv("USER_DATA_PATH")
	server.SetRoutes()

	database.Init()
	models.SetDefaultUsers()
	return server
}

func (server *Server) Run() {
	log.Println("Listening: " + server.Address + server.Port)
	server.Router.Run(server.Address + ":" + server.Port)
}

func (server *Server) SetRoutes() {
	if server.Router == nil {
		server.Router = gin.New()
	}
	// set a logger
	server.Router.Use(gin.Logger())
	// routes
	server.Router.POST("/signup", controllers.SignUp())
	server.Router.POST("/login", controllers.Login())

	authorized := server.Router.Group("/")

	authorized.Use(middlewares.AuthRequired())
	{
		authorized.POST("/add/users", controllers.AddUsers())
		authorized.DELETE("/delete/user/:id", controllers.DeleteUser())
		authorized.PATCH("/update/user/:id", controllers.UpdateUser())
		authorized.GET("/user/:id", controllers.GetUser())
	}
}
