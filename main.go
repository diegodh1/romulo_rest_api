package main

import (
	"log"
	handler "romulo/handler"

	"github.com/gin-gonic/gin"
)

//LoginUser struct login
type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	//connect to db
	var connection *handler.Config
	connection = new(handler.Config)
	connection.Init()
	_, err := connection.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola",
		})
	})

	r.POST("/login/user", func(c *gin.Context) {
		var userLogin LoginUser
		err := c.BindJSON(&userLogin)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"message": "Petición mal estructurada",
				"payload": nil,
			})
		case userLogin.Username == "" || userLogin.Password == "":
			c.JSON(400, gin.H{
				"message": "Por favor ingrese un usuario y/o contraseña válidos",
				"payload": nil,
			})
		default:
			c.JSON(200, gin.H{
				"message": "ok",
				"payload": "lol",
			})
		}
	})

	r.Run(":4000")
}
