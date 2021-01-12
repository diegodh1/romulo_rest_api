package main

import (
	"log"
	handler "romulo/handler"
	routes "romulo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//connect to db
	var connection *handler.Config
	connection = new(handler.Config)
	connection.Init()
	db, err := connection.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	r.POST("/user/login", routes.Login(db))
	r.POST("/user/create", routes.CreateUser(db))
	r.POST("/user/update", routes.UpdateUser(db))
	r.POST("/user/assignprofile", routes.AssignProfile(db))
	r.GET("/client/search/:name", routes.SearchClient(db))
	r.GET("/client/info/:nit", routes.GetInfoClient(db))
	r.Run(":4000")
}
