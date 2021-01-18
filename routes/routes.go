package routes

import (
	handler "romulo/handler"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//LoginUser struct login
type LoginUser struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

//USER ROUTES

//Login route
func Login(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var userLogin LoginUser
		err := c.BindJSON(&userLogin)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"message": "Petición mal estructurada",
				"payload": nil,
				"status":  400,
			})
		default:
			response := handler.LoginUser(userLogin.UserID, userLogin.Password, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

//CreateUser func
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user handler.User
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.CreateUser(&user.User, &user.Profiles, db)
			c.JSON(400, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

//UpdateUser func
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	var user handler.User
	fn := func(c *gin.Context) {
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil,
				"message": "petición mal estructurada",
				"status":  400,
			})
		default:
			response := handler.UpdateUser(&user.User, &user.Profiles, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return fn
}

//***************************CLIENT****************************

//SearchClient func
func SearchClient(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		name := c.Param("name")
		name = strings.ToUpper(strings.ReplaceAll(name, "%", ""))
		response := handler.SearchClient(name, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//SearchUser func
func SearchUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userID := c.Param("userID")
		userID = strings.ToUpper(strings.ReplaceAll(userID, "%", ""))
		response := handler.SearchUser(userID, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}

//GetInfoClient func
func GetInfoClient(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		nit := c.Param("nit")
		response := handler.GetPersonalInfo(nit, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return fn
}
